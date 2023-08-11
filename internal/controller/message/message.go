package message

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	v1 "llmfarm/api/message/v1"
	"llmfarm/internal/model/entity"
	botService "llmfarm/internal/service/bot"
	service "llmfarm/internal/service/message"
	"llmfarm/internal/service/session"
	"llmfarm/library/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type MessageController struct{}

var taskMap map[string]string = make(map[string]string)

func New() *MessageController {
	return &MessageController{}
}

func (s *MessageController) Evaluate(ctx context.Context, req *v1.EvaluateReq) (res *v1.EvaluateRes, err error) {
	err = service.New().Evaluate(req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (s *MessageController) Compare(ctx context.Context, req *v1.CompareReq) (res *v1.CompareRes, err error) {
	ress, err := service.New().Compare(req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

func (c *MessageController) List(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	ress, _ := service.New().List(req.ChatId, session.Context.GetUser(ctx).Id)
	// fmt.Println("message list", ress)
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

// 修改active
func (c *MessageController) ModifyActive(ctx context.Context, req *v1.UpdateMessageReq) (res *v1.UpdateMessageRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	ress := service.New().UpdateNodeActive(req, session.Context.GetUser(ctx).Id)
	// fmt.Println("message list", ress)
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

type GPTRequest struct {
	Message        string              `json:"message"`
	PrefixMessages []map[string]string `json:"prefixMessages"`
}

type GMLRequest struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

// "chainId": 1,
// "question": "中国"
type Chain struct {
	ChainId    int                 `json:"chainId"`
	Question   string              `json:"question"`
	Uuid       string              `json:"uuid"`
	Messages   []map[string]string `json:"messages"`
	OPENAI_KEY string              `json:"OPENAI_KEY"`
	MessageId  int                 `json:"messageId"`
	DataSetIds []string            `json:"setIds"`
}

var keyIndex int = 0
var count int = 0

// 提问
func (c *MessageController) Question(ctx context.Context, req *v1.ConversationReq) (res *v1.ConversationRes, err error) {
	var message = &entity.Message{}
	userId := session.Context.GetUser(ctx).Id
	taskMap[fmt.Sprintf("userId%d", userId)] = "stop"
	fmt.Println("ctx==========", ctx)
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	// 计算输入的question的token
	fmt.Println("3333333")

	// contentTokenCount, _ := service.New().GetContextToken(req.Content)
	contentTokenCount := len(req.Content)
	httpResponse := r.Response
	fmt.Println("4444444")

	questionRes, _ := service.New().AskQuestion(req, userId)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	httpResponse.Header().Set("Content-Type", "text/event-stream")
	httpResponse.Header().Set("Cache-Control", "no-cache")
	httpResponse.Header().Set("Connection", "keep-alive")
	versionType, _ := g.Cfg().Get(ctx, "versionType")
	fmt.Println("0000000")

	tokenRes, usageToken, nodes, err := service.New().ParseChain(req.BotId, userId, versionType.String())
	maxToken := c.getMaxToken(nodes)

	fmt.Println("111111111")
	tokenRes = false
	maxToken = 4000
	if tokenRes || contentTokenCount > maxToken {
		value := ""
		if tokenRes {
			value = g.I18n().Tf(r.Context(), "当前已使用token为,已超出token额度", usageToken)
		} else {
			value = g.I18n().T(r.Context(), "您的输入字数过长，请适当减少输入字数")
		}

		r := strings.NewReader(value)
		mess, _ := c.SaveMessage(value, req.BotId, questionRes.ChatId, questionRes.Id, userId)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			// Call your callback function here
			ch := scanner.Text()
			fmt.Println(string(ch))
			data := base64.StdEncoding.EncodeToString([]byte(string(ch)))
			fmt.Fprintf(httpResponse.Writer, "data: %s\n\n", data)
			httpResponse.Writer.Flush()
		}

		fmt.Fprintf(httpResponse.Writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(questionRes.Id)), strconv.Itoa(int(mess.Id)), strconv.Itoa(int(questionRes.ChatId)))
		httpResponse.Writer.Flush()

		return
	}
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	//获取botId 关联的 ChainId
	bot, _ := botService.New().Detail(req.BotId, userId)

	//获取bot关联的文件Id
	// fileIds, err := botService.New().GetFileIds(req.BotId, userId)
	// if err != nil {
	// 	return nil, err
	// }
	// newFileIds := []string{}
	// for _, fileId := range fileIds {
	// 	newFileIds = append(newFileIds, gconv.String(fileId))
	// }
	dataSetIds, err := botService.New().GetBotDataSetIds(req.BotId)
	newDataSetIds := []string{}
	fmt.Println("dataSetIds======", dataSetIds)
	for _, dataSetId := range dataSetIds {
		fmt.Println("dataSetId======", dataSetId)
		newDataSetIds = append(newDataSetIds, gconv.String(dataSetId))
	}

	collectionName, _ := g.Cfg().Get(ctx, "api.collectionName")
	var resp *http.Response
	var data []byte
	contextParams := []map[string]string{}
	if req.ContextType == "连续对话" {
		contextParams, err = service.New().GetContext(req.ParentId, req.ChatId, questionRes.UserId, maxToken-contentTokenCount)
		if err != nil {
			fmt.Println("上下文获取失败", err)
			return nil, err
		}
	}
	message.Content = ""
	message.BotId = req.BotId
	message.ChatId = questionRes.ChatId
	message.ParentId = questionRes.Id
	//保存message
	answerRes, _ := service.New().SetAnswer(message, userId)

	request := &Chain{
		ChainId:    int(bot.ChainId),
		Question:   req.Content,
		Uuid:       collectionName.String(),
		Messages:   contextParams,
		MessageId:  int(answerRes.Id),
		OPENAI_KEY: "",
		DataSetIds: newDataSetIds,
	}
	data, _ = json.Marshal(request)
	fmt.Println(string(data))

	//调用ChainAPI
	resp = service.New().Chain(ctx, data)
	//设置keyCount
	keyIndex++
	if count == keyIndex {
		keyIndex = 0
	}

	message.Content = responseToBrowser(resp, httpResponse.Writer, questionRes.Id, answerRes.Id, message.ChatId, userId, answerRes, req.BotId)
	answerRes.Content = message.Content
	answerRes.BotId = req.BotId
	data, _ = json.Marshal(g.Map{"content": req.Content})
	if taskMap[fmt.Sprintf("userId%d", userId)] == "end" {
		fmt.Println("提问回答会话正常保存========")
		service.New().Update(answerRes, userId)
	}
	if resp != nil {
		// 读取Header
		UUID := resp.Header.Get("Chain-Execution-UUID")
		fmt.Println("UUID==============", UUID)
		service.New().GetTokens(ctx, answerRes.Id, userId, UUID)
	}

	fmt.Fprintf(httpResponse.Writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(questionRes.Id)), strconv.Itoa(int(answerRes.Id)), strconv.Itoa(int(questionRes.ChatId)))
	httpResponse.Writer.Flush()
	return
}

// 获取当前chain的最大token数
func (c *MessageController) getMaxToken(nodes []map[string]interface{}) int {
	maxToken := 15000

	// 获取所有llm模型，做个map
	var models []entity.Model
	g.Model("model").Where("deleteTime is null").Scan(&models)
	llmMap := map[string]int{}
	for _, model := range models {
		llmMap[model.Tag] = model.MaxTokenLimit
	}

	// 找到chain里面maxTokenLimit最小的模型
	for _, node := range nodes {
		tag := getModelTag(node)
		fmt.Println("tag:", tag)
		if llmMap[tag] != 0 {
			if llmMap[tag] < maxToken {
				maxToken = llmMap[tag]
			}
		}
	}

	// 获取llm模型的最大token数
	return maxToken
}

func getModelTag(node map[string]interface{}) string {
	data := node["data"].(map[string]interface{})
	nodeName := data["type"].(string)
	if nodeName == "ChatCompletion" {
		value := data["value"].(map[string]interface{})
		return value["modelName"].(string)
	} else {
		return nodeName
	}
}

// 保存Message
func (c *MessageController) SaveMessage(value string, botId uint, chatId uint, parentId uint, userId uint) (*entity.Message, error) {
	var message = &entity.Message{}
	message.Content = value
	message.BotId = botId
	message.ChatId = chatId
	message.ParentId = parentId
	answerRes, _ := service.New().SetAnswer(message, userId)
	if answerRes != nil {
		return answerRes, nil
	}
	return answerRes, nil
}

// 获取当日已使用token额度
func (c *MessageController) GetUserDailyUsage(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	userId := session.Context.GetUser(ctx).Id

	usageToken, err := service.New().GetUserDailyUsage(userId)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", usageToken)
	return
}

// 流式返回信息
func ReturnPrompt(writer *ghttp.ResponseWriter, r *strings.Reader, questionId uint, answerId uint, chatId uint) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		// Call your callback function here
		ch := scanner.Text()
		fmt.Println(string(ch))
		data := base64.StdEncoding.EncodeToString([]byte(string(ch)))
		fmt.Fprintf(writer, "data: %s\n\n", data)
		writer.Flush()
	}
	fmt.Fprintf(writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(questionId)), strconv.Itoa(int(answerId)), strconv.Itoa(int(chatId)))
	writer.Flush()
	return
}

func responseToBrowser(resp *http.Response, writer *ghttp.ResponseWriter, questionId uint, answerId uint, chatId uint, userId uint, answerRes *entity.Message, botId uint) string {
	if resp == nil {
		info := g.I18n().T(context.TODO(), "很抱歉给您带来不便，由于当前访问量过大，我们的系统无法满足所有用户的请求，建议您稍后重新提问。")
		r := strings.NewReader(info)
		ReturnPrompt(writer, r, questionId, answerId, chatId)
		return info
	}
	fmt.Println("resp.StatusCode===============", resp.StatusCode)
	if resp.StatusCode != 200 {
		info := fmt.Sprintf("请求异常,状态码：%d", resp.StatusCode)
		r := strings.NewReader(info)
		ReturnPrompt(writer, r, questionId, answerId, chatId)
		return info
	}
	defer resp.Body.Close()
	var content = ""
	// buf := make([]byte, 1024)
	taskMap[fmt.Sprintf("userId%d", userId)] = "begin"

	scanner := bufio.NewScanner(resp.Body)

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		if taskMap[fmt.Sprintf("userId%d", userId)] == "stop" && taskMap[fmt.Sprintf("userId%dMessageId%d", userId, answerId)] != "stops" {
			fmt.Println("停止提问了")
			fmt.Fprintf(writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(questionId)), strconv.Itoa(int(answerId)), strconv.Itoa(int(chatId)))
			writer.Flush()
			taskMap[fmt.Sprintf("userId%dMessageId%d", userId, answerId)] = "stops"
			answerRes.Content = content
			answerRes.BotId = botId
			service.New().Update(answerRes, userId)
			fmt.Println("停止会话保存=========:", content)

		} else if taskMap[fmt.Sprintf("userId%d", userId)] == "begin" && taskMap[fmt.Sprintf("userId%dMessageId%d", userId, answerId)] != "stops" {

			ch := scanner.Text()
			fmt.Println("=======", string(ch))
			content += string(ch)
			data := base64.StdEncoding.EncodeToString([]byte(string(ch)))
			fmt.Fprintf(writer, "data: %s\n\n", data)
			writer.Flush()
		}
	}

	fmt.Println("问答结束", taskMap[fmt.Sprintf("userId%d", userId)])
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Println("Error reading response:", err)
		}
	}
	taskMap[fmt.Sprintf("userId%d", userId)] = "end"
	return content
}

// 重新回答
func (c *MessageController) ReAnswer(ctx context.Context, req *v1.ReconsiderAnswerReq) (res *v1.ReconsiderAnswerRes, err error) {
	userId := session.Context.GetUser(ctx).Id
	//1 根据parentId获取message
	message, _ := service.New().FindMessage(req.ParentId, userId)
	//2 根据message获取content
	context := message.Content
	fmt.Println("context1:", context)
	// 计算输入的question的token
	contentTokenCount, _ := service.New().GetContextToken(context)
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	httpResponse := r.Response
	//questionRes, _ := service.New().AskQuestion(conversationReq, userId)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	httpResponse.Header().Set("Content-Type", "text/event-stream")
	httpResponse.Header().Set("Cache-Control", "no-cache")
	httpResponse.Header().Set("Connection", "keep-alive")
	versionType, _ := g.Cfg().Get(ctx, "versionType")
	tokenRes, usageToken, nodes, err := service.New().ParseChain(req.BotId, userId, versionType.String())
	maxToken := c.getMaxToken(nodes)
	if tokenRes || contentTokenCount > maxToken {
		value := ""
		if tokenRes {
			// "当前已使用token为"
			value = g.I18n().Tf(r.Context(), "当前已使用token为,已超出token额度", usageToken)
		} else {
			value = g.I18n().T(r.Context(), "您的输入字数过长，请适当减少输入字数")
			// value = "您的输入字数过长，请适当减少输入字数"
		}
		r := strings.NewReader(value)
		mess, _ := c.SaveMessage(value, req.BotId, message.ChatId, message.Id, userId)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			// Call your callback function here
			ch := scanner.Text()
			fmt.Println(string(ch))
			data := base64.StdEncoding.EncodeToString([]byte(string(ch)))
			fmt.Fprintf(httpResponse.Writer, "data: %s\n\n", data)
			httpResponse.Writer.Flush()
		}

		fmt.Fprintf(httpResponse.Writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(message.ParentId)), strconv.Itoa(int(mess.Id)), strconv.Itoa(int(message.ChatId)))
		httpResponse.Writer.Flush()

		return
	}
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	//获取botId 关联的 ChainId
	bot, _ := botService.New().Detail(req.BotId, userId)

	dataSetIds, err := botService.New().GetBotDataSetIds(req.BotId)
	newDataSetIds := []string{}
	for _, dataSetId := range dataSetIds {
		newDataSetIds = append(newDataSetIds, gconv.String(dataSetId))
	}
	collectionName, _ := g.Cfg().Get(ctx, "api.collectionName")
	var resp *http.Response
	var data []byte
	contextParams := []map[string]string{}
	if req.ContextType == "连续对话" {
		contextParams, err = service.New().GetContext(message.ParentId, req.ChatId, message.UserId, maxToken-contentTokenCount)
		if err != nil {
			fmt.Println("上下文获取失败", err)
			return nil, err
		}
	}
	fmt.Println("contextParamsOld", contextParams)
	message.Content = ""
	message.BotId = req.BotId
	message.ParentId = message.Id
	//保存message
	answerRes, _ := service.New().SetAnswer(message, userId)

	request := &Chain{
		ChainId:    int(bot.ChainId),
		Question:   context,
		Uuid:       collectionName.String(),
		Messages:   contextParams,
		MessageId:  int(answerRes.Id),
		OPENAI_KEY: "",
		DataSetIds: newDataSetIds,
	}
	data, _ = json.Marshal(request)

	//调用ChainAPI
	resp = service.New().Chain(ctx, data)
	//设置keyCount
	keyIndex++
	if count == keyIndex {
		keyIndex = 0
	}

	message.Content = responseToBrowser(resp, httpResponse.Writer, message.Id, answerRes.Id, message.ChatId, userId, answerRes, req.BotId)
	answerRes.Content = message.Content
	answerRes.BotId = req.BotId
	//data, _ = json.Marshal(g.Map{"content": context})

	if taskMap[fmt.Sprintf("userId%d", userId)] == "end" {
		fmt.Println("重新回答会话正常保存========")
		service.New().Update(answerRes, userId)
	}

	if resp != nil {
		UUID := resp.Header.Get("Chain-Execution-UUID")
		service.New().GetTokens(ctx, answerRes.Id, userId, UUID)
	}
	// 读取Header

	fmt.Fprintf(httpResponse.Writer, "data:__47b6a0a0ebd4__{\"questionId\":%s,\"answerId\":%s,\"chatId\":%s}\n\n", strconv.Itoa(int(message.Id)), strconv.Itoa(int(answerRes.Id)), strconv.Itoa(int(message.ChatId)))
	httpResponse.Writer.Flush()
	return
}

// 停止回答
func (c *MessageController) StopAnswer(ctx context.Context, req *v1.StopAnswerReq) (res *v1.StopAnswerRes, err error) {
	userId := session.Context.GetUser(ctx).Id
	fmt.Println(req.ChatId, "======================", taskMap[fmt.Sprintf("userId%d", userId)], "=====================================")
	taskMap[fmt.Sprintf("userId%d", userId)] = "stop"
	fmt.Println(req.ChatId, "======================", taskMap[fmt.Sprintf("userId%d", userId)], "=====================================")
	return
}
