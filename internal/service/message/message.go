package message

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	chatv1 "llmfarm/api/chat/v1"
	v1 "llmfarm/api/message/v1"
	"llmfarm/internal/model/entity"
	botService "llmfarm/internal/service/bot"
	chainService "llmfarm/internal/service/chain"
	service "llmfarm/internal/service/chat"
	tokenService "llmfarm/internal/service/tokenUsageLimit"
	userService "llmfarm/internal/service/user"
	"net/http"

	"github.com/pkoukk/tiktoken-go"

	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type MessageService struct {
}

func New() *MessageService {
	return &MessageService{}
}

func (s *MessageService) Evaluate(req *v1.EvaluateReq) error {
	var message *entity.Message
	m := g.Model("message")
	r, err := m.Where("id=?", req.MessageId).One()
	if err != nil {
		fmt.Println("err", err)
		return errors.New("消息不存在")
	}
	r.Struct(&message)
	if message == nil {
		return errors.New("消息不存在")
	}
	_, err = m.Data(g.Map{"mark": req.Mark, "feedback": req.Feedback}).Where("id=?", message.Id).Update()
	if err != nil {
		fmt.Println("err", err)
		return errors.New("评价失败")
	}
	return nil
}

func (s *MessageService) Compare(req *v1.CompareReq) (*entity.CompareRecord, error) {
	var compareRecord = new(entity.CompareRecord)
	compareRecord.DefaultCreateTime()
	compareRecord.DefaultUpdateTime()
	compareRecord.ParentId = req.ParentId
	compareRecord.SourceMessageId = req.SourceMessageId
	compareRecord.TargetMessageId = req.TargetMessageId
	compareRecord.Result = req.Result
	r, err := g.Model("compare_record").Insert(&compareRecord)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("对比失败")
	}
	compareRecordId, err := r.LastInsertId()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("对比失败")
	}
	compareRecord.Id = uint(compareRecordId)
	return compareRecord, nil
}

func (s *MessageService) List(chatId uint, userId uint) ([]*v1.Message, error) {
	var messages []*v1.Message
	// 	SELECT m.*,b.botName,b.icon from message m
	// LEFT JOIN bot b on b.id = m.botId
	res, err := g.Model("message m").Fields("m.*,b.botName,b.icon").LeftJoin("bot b", "b.id = m.botId").Where("m.chatId = ? AND m.userId = ? AND m.deleteTime IS null", chatId, userId).All()
	if err != nil {
		return nil, errors.New("chatId不存在")
	}
	res.Structs(&messages)
	fmt.Println("List", messages)
	return messages, nil
}

func (s *MessageService) FindMessage(messageId uint, userId uint) (*entity.Message, error) {
	var message *entity.Message
	res, err := g.Model("message").Where("id = ? AND userId = ? AND deleteTime IS null", messageId, userId).One()
	if err != nil {
		fmt.Println("查询message失败", err)
		return nil, err
	}
	res.Struct(&message)
	return message, nil
}

func (s *MessageService) Create(req *entity.Message, userId uint) (*entity.Message, error) {
	fmt.Println("chatId", req, req.ChatId)
	message := entity.Message{
		ChatId:      req.ChatId,
		Content:     req.Content,
		ParentId:    req.ParentId,
		UserId:      userId,
		Mark:        "1",
		Type:        req.Type,
		ContentType: req.ContentType,
		Active:      req.Active,
		BotId:       req.BotId,
	}
	message.DefaultCreateTime()
	message.DefaultUpdateTime()
	res, err := g.Model("message").Insert(&message)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	messageId, err := res.LastInsertId()
	message.Id = uint(messageId)
	return &message, err
}

func (s *MessageService) Update(req *entity.Message, userId uint) error {
	var params = g.Map{}
	if req.Mark != "" {
		params["mark"] = req.Mark
	}
	if req.Active != 0 {
		params["active"] = req.Active
	}
	if req.Content != "" {
		params["content"] = req.Content
	}
	// if req.UsageTotalTokens > 0 {
	// 	params["usageTotalTokens"] = req.UsageTotalTokens
	// }
	params["updateTime"] = gtime.Now()
	_, err := g.Model("message").Data(params).Where("id = ? and userId = ? ", req.Id, userId).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) Delete(id uint, userId uint) error {
	_, err := g.Model("message").Where("id = ? and userId = ? ", id, userId).Delete()
	if err != nil {
		return err
	}
	return nil
}

// 提问，重新提问
func (s *MessageService) AskQuestion(req *v1.ConversationReq, userId uint) (*entity.Message, error) {
	var userMessage = new(entity.Message)
	//判断chatId是否存在	不存在则创建
	if req.ChatId == 0 {
		var chat = chatv1.CreateChatReq{
			Name:        gstr.SubStrRune(req.Content, 0, 40),
			Type:        "1",
			ContextType: req.ContextType,
			IsTemp:      req.IsTemp,
		}
		newChat, err := service.New().Create(&chat, userId)
		if err != nil {
			return nil, errors.New("新建提问chat失败")
		}
		req.ChatId = uint(newChat.Id)
	}
	userMessage.ChatId = req.ChatId
	userMessage.Content = req.Content
	userMessage.ParentId = req.ParentId
	userMessage.BotId = req.BotId
	userMessage.UserId = userId
	userMessage.Mark = "1"
	userMessage.Type = "2"
	userMessage.ContentType = "1"
	userMessage.Active = 1
	userMessage.DefaultCreateTime()
	userMessage.DefaultUpdateTime()
	fmt.Println("chatId", req.ChatId)
	res, err := s.Create(userMessage, userId)
	if err != nil {
		return nil, errors.New("新建提问失败")
	}
	return res, nil
}

func (s *MessageService) SetAnswer(req *entity.Message, userId uint) (*entity.Message, error) {
	var userMessage = new(entity.Message)

	userMessage.ChatId = req.ChatId
	userMessage.Content = req.Content
	userMessage.ParentId = req.ParentId
	userMessage.UserId = userId
	userMessage.Mark = "1"
	userMessage.Type = "1"
	userMessage.ContentType = "1"
	userMessage.Active = 1
	userMessage.BotId = req.BotId
	userMessage.DefaultCreateTime()
	userMessage.DefaultUpdateTime()
	fmt.Println("SetAnswer", req.ChatId)

	res, err := s.Create(userMessage, userId)
	return res, err
}

// 修改节点active
func (s *MessageService) UpdateNodeActive(req *v1.UpdateMessageReq, userId uint) error {
	var updateRes *entity.Message
	fmt.Println("req", req.Id)
	if res, err := s.FindMessage(req.Id, userId); err == nil {
		updateRes = res
		fmt.Println("updateRes123", res, err, updateRes)
	}
	unShowMessage := g.Map{"active": 0, "updateTime": gtime.Now()}
	if _, err := g.Model("message").Data(&unShowMessage).Where("parentId = ? AND userId = ?", updateRes.ParentId, userId).Update(); err != nil {
		return errors.New("全部更新失败")
	}
	showMessage := g.Map{"active": 0, "updateTime": gtime.Now()}
	if _, err := g.Model("message").Data(&showMessage).Where("id = ? ", req.Id).Update(); err != nil {
		return errors.New("更新失败")
	}
	return nil
}

// 获取上下文
func (s *MessageService) GetContext(parentId uint, chatId uint, userId uint, maxToken int) ([]map[string]string, error) {
	fmt.Println("111111111", maxToken)
	messages, err := s.List(chatId, userId)
	if err != nil {
		return nil, err
	}

	messageMap := make(map[uint]*v1.Message)
	for _, message := range messages {
		messageMap[message.Id] = message
	}

	var messageTokens []entity.MessageModelUsageTokens
	//获取ID数组
	var ids []uint
	for _, v := range messages {
		ids = append(ids, gconv.Uint(v.Id))
	}
	//获取所有的token
	res, err := g.Model("messageModelUsageTokens").Where("messageId in (?)", ids).OrderAsc("createTime").All()
	res.Structs(&messageTokens)

	messageTokensMap := make(map[uint]int)
	for _, v := range messageTokens {
		messageTokensMap[v.Id] = int(v.UsageTotalTokens)
	}

	// a, _ := json.Marshal(messageMap)
	// fmt.Println("00000", string(a))
	tokenCount := 0
	var params []map[string]string
	var newParams []map[string]string
	newParentId := parentId
	for messageMap[newParentId] != nil && tokenCount < maxToken {
		contextToken, _ := s.GetContextToken(messageMap[newParentId].Content)
		tokenCount += contextToken
		if tokenCount > maxToken {
			continue
		}
		itemMap := make(map[string]string)
		role := "user"
		if messageMap[newParentId].Type == "BOT" {
			role = "assistant"
		}
		itemMap["role"] = role
		itemMap["content"] = messageMap[newParentId].Content
		params = append(params, itemMap)
		newParentId = messageMap[newParentId].ParentId
	}

	for i := len(params) - 1; i >= 0; i-- {
		newParams = append(newParams, params[i])
	}

	// fmt.Println("1111111", string(a))

	return newParams, err
}

func (s *MessageService) Chain(ctx context.Context, data []byte) *http.Response {
	r := g.RequestFromCtx(ctx)
	ip := r.Header.Get("x-real-ip")
	if ip == "" {
		ip = r.Header.Get("x-forwarded-for")
	}
	fmt.Printf("用户请求 IP:%s\n", ip)

	chainApiDomain, _ := g.Cfg().Get(ctx, "api.chainApiDomain")
	url := fmt.Sprintf("%s/api/chainflow/run", chainApiDomain)
	fmt.Println("url", url)
	//url := "https://chain-dev.wuwei-prod.wudaima.com/api/chainflow/run"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-real-ip", ip)
	fmt.Println("request Header: ", request.Header)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return nil
	}
	return resp
}

func (s *MessageService) ChatGLM(ctx context.Context, data []byte) *http.Response {
	//chatGLMApiDomain, _ := g.Cfg().Get(ctx, "api.chatGLMApiDomain")
	//url := fmt.Sprintf("%s/v1/chat/completions", chatGLMApiDomain)
	url := "http://39.97.107.228:7001/v1/chat/completions"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer token1")

	fmt.Println("request Header: ", request.Header)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return nil
	}
	return resp
}

type Response struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func (s *MessageService) GetTokens(ctx context.Context, messageId, userId uint, uuid string) error {
	chainApiDomain, _ := g.Cfg().Get(ctx, "api.chainApiDomain")
	url := fmt.Sprintf("%s/api/chainflow/log?uuid=%s", chainApiDomain, uuid)
	fmt.Println("url==================", url)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return err
	}
	defer resp.Body.Close()
	// 获取 tokenConsumed
	body, _ := ioutil.ReadAll(resp.Body)
	var resultBody map[string]interface{}
	gconv.Struct(body, &resultBody)
	a, _ := json.Marshal(resultBody)
	var result1 Response
	err = json.Unmarshal(a, &result1)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
	}
	if result1.Data["tokenConsumed"] == nil {
		fmt.Println("result1.Data", result1.Data)
		fmt.Println("解析tokenConsumed失败:", err)
		return nil
	}
	fmt.Println("tokenConsumed:", result1.Data["tokenConsumed"])

	// 判断用户信息

	//fmt.Println("Code值:", result1.Data.(map[string]interface{})["tokenConsumed"])
	fmt.Println("Msg值:", result1.Msg)
	if err != nil {
		return err
	}
	return nil

	// 	dataStr := string(data)
	// 	return uint(gstr.LenRune(dataStr))
}

// 单独获取一段文本的token
func (s *MessageService) GetContextToken(content string) (int, error) {
	encoding := "cl100k_base"
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return 0, err
	}
	token := tke.Encode(content, nil, nil)
	return len(token), nil
}

// 解析chain
func (s *MessageService) ParseChain(botId uint, userId uint, versionType string) (bool, uint, []map[string]interface{}, error) {
	// 先根据botId获取chainId
	bot, err := botService.New().Detail(botId, userId)
	if err != nil {
		return false, 0, nil, err
	}
	chainId := bot.ChainId
	// 根据chainId获取modelId
	chain, err := chainService.New().Detail(chainId, userId)
	if err != nil {
		return false, 0, nil, err
	}
	nodeInfo := chain.Chain.Nodes

	// 判断用户信息
	userInfo, err := userService.New().UpdateUserById(userId)
	if err != nil {
		fmt.Println("获取用户信息失败", err)
		return false, 0, nil, err
	}
	usageToken := uint(0)
	dailyTokens := uint(0)
	//判断用户身份
	if userInfo.Is_membership == 0 || versionType == "enterprise" {
		//非会员 或 企业用户
		//获取已使用量
		usageToken, _ = s.GetUserDailyUsage(userId)
		//获取token总数
		dailyTokens, _ = tokenService.New().GetUserdailyTokens(userId)
	} else if userInfo.Is_membership == 1 {
		//会员用户
		//获取已使用量 和 总量
		if err != nil {
			return false, 0, nil, err
		}
	}
	//如果token总数大于等于已使用量 则返回true
	if usageToken >= dailyTokens {
		return true, usageToken, nodeInfo, nil
	}
	return false, 0, nodeInfo, nil
}

// 获取用户当日token已使用量
func (s *MessageService) GetUserDailyUsage(userId uint) (uint, error) {

	type MessageResult struct {
		UsageTotalTokens uint
	}
	var messages []*MessageResult
	err := g.Model("messageModelUsageTokens", "mtu").LeftJoin("message m", "mtu.messageId = m.id").Where("m.userId = ? ", userId).Fields("mtu.usageTotalTokens").Scan(&messages)
	//err := g.Model("message").Where("userId = ? AND botId = ? AND DATE(DATE_ADD(createTime, INTERVAL 8 HOUR))=CURDATE() AND deleteTime IS null", userId, botId).Fields("usageTotalTokens").Scan(&messages)
	if err != nil {
		return 0, errors.New("获取message列表失败")
	}

	var total uint
	total = 0
	for _, message := range messages {
		total = total + message.UsageTotalTokens
	}
	return total, nil
}
