package chain

import (
	"errors"
	botV1 "llmfarm/api/bot/v1"
	v1 "llmfarm/api/chain/v1"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ChainService struct{}

func New() *ChainService {
	return &ChainService{}
}

// 获取chain的类型
func (s *ChainService) getChainType(nodes []map[string]interface{}) (int, error) {
	chainMap := gmap.New()
	chainMap.Set("OpenAIEmbedding", "1")
	chainMap.Set("VectorSearch", "1")
	chainMap.Set("SQLInsert", "1")
	chainMap.Set("SQLRun", "1")
	chainMap.Set("SQLTableSchema", "1")
	chainMap.Set("AMapIP", "1")
	chainMap.Set("AMapWeather", "1")

	count := 0
	for _, node := range nodes {
		nodeType := node["data"].(map[string]interface{})["type"].(string)
		if chainMap.Get(nodeType) != nil {
			return 3, nil
		}
		if nodeType == "PromptTemplate" {
			count++
		}
	}
	if count > 1 {
		return 2, nil
	}

	if count == 1 {
		return 1, nil
	}

	return 1, nil
}

// copy
func (s *ChainService) CopyChain(chain *entity.Chain, userId uint) (*entity.Chain, error) {
	var chainNew = new(entity.Chain)
	chainNew.Edges = chain.Edges
	chainNew.ChainName = chain.ChainName
	chainNew.UserId = userId
	chainNew.UserType = chain.UserType
	chainNew.Nodes = chain.Nodes
	chainNew.ChainLevel = chain.ChainLevel
	chainNew.UseKnowledge = chain.UseKnowledge
	chainNew.DefaultCreateTime()
	chainNew.DefaultUpdateTime()
	res, err := g.Model("chain").Insert(&chainNew)
	if err != nil {
		return nil, errors.New("流程复制失败")
	}
	chainId, err := res.LastInsertId()
	chainNew.Id = uint(chainId)
	return chainNew, err
}

func (s *ChainService) Create(req *v1.CreateChainReq, userId uint) (*entity.Chain, error) {
	var chain = new(entity.Chain)
	chain.ChainName = req.Name
	chain.UserId = userId
	chain.Nodes = req.Nodes
	chain.Edges = req.Edges
	chain.UserType = "user"
	useKnowledge := 0
	for _, node := range req.Nodes {
		data := node["data"].(map[string]interface{})
		if data["type"] != nil {
			nodeType := data["type"].(string)
			if nodeType == "VectorSearch" {
				useKnowledge = 1
			}
		}

	}

	chain.UseKnowledge = uint(useKnowledge)
	chain.DefaultCreateTime()
	chain.DefaultUpdateTime()
	res, err := g.Model("chain").Insert(&chain)
	if err != nil {
		return nil, errors.New("创建失败")
	}
	chainId, err := res.LastInsertId()
	chain.Id = uint(chainId)

	return chain, err
}

func (s *ChainService) createBot(req *botV1.CreateBotReq, userId uint) (*entity.Bot, error) {
	var bot *entity.Bot = new(entity.Bot)
	bot.UserId = userId
	bot.BotName = req.Name
	bot.ChainId = req.ChainId
	bot.UserType = "user"
	bot.CreateTime = gtime.Now()
	bot.UpdateTime = gtime.Now()

	res, err := g.Model("bot").Insert(&bot)
	if err != nil {
		return nil, errors.New("创建失败")
	}
	botId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("创建失败")
	}
	bot.Id = uint(botId)

	return bot, err
}

func (s *ChainService) UpdateChainName(req *v1.UpdateChainNameReq) error {
	var chain = new(entity.Chain)
	r, err := g.Model("chain").Where("id = ? and deleteTime is null", req.Id).One()
	if err != nil {
		return errors.New("流程不存在")
	}
	r.Struct(&chain)

	if chain.CreateTime == nil {
		return errors.New("流程不存在")
	}
	_, err = g.Model("chain").Data(chain).Where("id = ?", req.Id).Update(g.Map{
		"chainName":  req.Name,
		"updateTime": gtime.Now(),
	})
	if err != nil {
		return errors.New("修改失败")
	}
	return err
}

func (s *ChainService) Update(req *v1.ModifyChainReq, userId uint) error {
	var chain = new(entity.Chain)
	r, err := g.Model("chain").Where("id = ? and deleteTime is null", req.Id).One()
	//useKnowledge
	//解析nodes
	useKnowledge := 0
	for _, node := range req.Nodes {
		data := node["data"].(map[string]interface{})
		if data["type"] != nil {
			nodeType := data["type"].(string)
			if nodeType == "VectorSearch" {
				useKnowledge = 1
			}
		}

	}
	if err != nil {
		return errors.New("流程不存在")
	}
	r.Struct(&chain)

	if chain.CreateTime == nil {
		return errors.New("流程不存在")
	}

	if chain.UserType == "system" {
		return errors.New("系统chain不能编辑")
	}
	_, err = g.Model("chain").Data(g.Map{
		"nodes":        req.Nodes,
		"edges":        req.Edges,
		"updateTime":   gtime.Now(),
		"useKnowledge": useKnowledge,
	}).Where("id = ?", req.Id).Update()
	if err != nil {
		return errors.New("修改失败")
	}

	// 如果没有bot，创建bot
	count, _ := g.Model("bot").Count("chainId = ? and deleteTime is null", chain.Id)
	// 创建chain自动创建相关联的bot
	if count == 0 {
		var botParam = new(botV1.CreateBotReq)
		botParam.IsPublic = false
		botParam.Name = chain.ChainName
		botParam.ChainId = chain.Id

		s.createBot(botParam, userId)
	}

	return err
}

func (s *ChainService) Delete(req *v1.DeleteChainReq) error {
	var chain = new(entity.Chain)
	r, err := g.Model("chain").Where("id = ? and deleteTime is null", req.Id).One()
	if err != nil {
		return errors.New("流程不存在")
	}
	r.Struct(&chain)

	if chain.CreateTime == nil {
		return errors.New("流程不存在")
	}

	if chain.UserType == "system" {
		return errors.New("系统chain不能删除")
	}

	_, err = g.Model("chain").Data(g.Map{"deleteTime": gtime.Now()}).Where("id", req.Id).Update()
	if err != nil {
		return errors.New("删除失败")
	}

	g.Model("bot").Data(g.Map{"deleteTime": gtime.Now()}).Where("chainId = ? and deleteTime is null", req.Id).Update()
	return err
}

type Chain struct {
	entity.Chain
	ShareOpen      bool   `json:"shareOpen" default:"false" `
	UserPermission string `json:"userPermission" dc:"用户权限"`
}

func (s *ChainService) List(req *v1.ListChainReq, userId uint) ([]Chain, error) {
	//查询关联表 chainBotUserRelation
	// var chainBotUserRelation []entity.ChainBotUserRelation
	// chainBotUserRelationR, err := g.Model("chain_bot_user_relation").Where("userId = ? and relationType = ?", userId, "chain").All()
	// if err != nil {
	// 	return nil, errors.New("查询失败")
	// }
	// chainBotUserRelationR.Structs(&chainBotUserRelation)

	var chainIds []int
	// for _, relation := range chainBotUserRelation {
	// 	chainIds = append(chainIds, relation.RelationId)
	// }

	var chains []Chain
	//mBot := g.Model("bot").Where("deleteTime is null and userType = ? and userId=?", "user", userId)
	systemChainsList := g.Model("chain").Where("deleteTime is null and chainName like ? and userType = 'user' ", "%"+req.Word+"%")
	if len(chainIds) > 0 {
		systemChainsList = systemChainsList.Where("( userId = ? OR id in (?) )", userId, chainIds)
	} else {
		systemChainsList = systemChainsList.Where("userId = ?", userId)
	}
	r, err := systemChainsList.OrderDesc("createTime").All()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	r.Structs(&chains)
	//记录我创建的chainId
	var myChainIds []int
	for i, chain := range chains {
		//判断userId 如果是创建者，返回creator
		if chain.UserId == userId {

			chains[i].UserPermission = "creator"

			myChainIds = append(myChainIds, int(chain.Id))
		} else {
			// 判断 Id 在不在 chainIds 里面
			for _, chainId := range chainIds {
				if int(chain.Id) == chainId {

					chains[i].UserPermission = "share"
					break
				}
			}
		}
	}

	//查询一下我的chain 看看那些开启分享了
	// if len(myChainIds) > 0 {
	// 	//如果chain_bot_user_relation 中有chainId 则说明开启分享了
	// 	var shareChainIds []int
	// 	shareChainIdsR, err := g.Model("chain_bot_user_relation").Where("relationId in (?) and relationType = ? and deleteTime is null", myChainIds, "chain").All()
	// 	if err != nil {
	// 		return nil, errors.New("查询失败")
	// 	}
	// 	var shareChainBotUserRelation []entity.ChainBotUserRelation
	// 	shareChainIdsR.Structs(&shareChainBotUserRelation)
	// 	for _, relation := range shareChainBotUserRelation {
	// 		shareChainIds = append(shareChainIds, relation.RelationId)
	// 	}
	// 	for i, chain := range chains {
	// 		for _, shareChainId := range shareChainIds {
	// 			if int(chain.Id) == shareChainId {
	// 				chains[i].ShareOpen = true
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	return chains, err
}

func (s *ChainService) Detail(chainId uint, userId uint) (*v1.ChainInfo, error) {
	var chain = new(v1.ChainInfo)
	r, err := g.Model("chain").
		InnerJoin("user u", "u.id = chain.userId").
		Fields("chain.*,u.name as userName,u.phone as userPhone").
		Where("chain.id = ? and chain.deleteTime is null", chainId).One()
	if err != nil {
		return nil, errors.New("流程不存在")
	}
	r.Struct(&chain)

	if chain.CreateTime == nil {
		return nil, errors.New("流程不存在")
	}

	if chain.UserId == userId {
		chain.UserPermission = "creator"
	} else {
		//去关联表里查询
		var chainBotUserRelation = new(entity.ChainBotUserRelation)
		r, err := g.Model("chain_bot_user_relation").Where("relationId = ? and userId = ? and relationType = ? and deleteTime is null", chainId, userId, "chain").One()
		if err != nil {
			return nil, errors.New("查询失败")
		}
		r.Struct(&chainBotUserRelation)
		chain.UserPermission = chainBotUserRelation.UserPermission
	}

	return chain, err
}
