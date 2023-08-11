package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ListChainReq struct {
	g.Meta `path:"/chain/list" tags:"提示链相关接口" method:"get" summary:"查询提示链列表"`
	Word   string `json:"word" dc:"关键词"`
}

type ListChainRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.Chain `json:"data" dc:"数据体"`
}

type DetailChainReq struct {
	g.Meta `path:"/chain/detail" tags:"提示链相关接口" method:"get" summary:"查询提示链详情"`
	Id     uint `json:"id" dc:"提示链id" v:"required"`
}

type ChainInfo struct {
	entity.Chain
	UserName       string `json:"userName" dc:"用户姓名"`
	UserPhone      string `json:"userPhone" dc:"用户手机号"`
	UserPermission string `json:"userPermission" dc:"用户权限"`
}
type DetailChainRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data ChainInfo `json:"data" dc:"数据体"`
}

type DeleteChainReq struct {
	g.Meta `path:"/chain" tags:"提示链相关接口" method:"delete" summary:"删除提示链"`
	Id     uint `json:"id" dc:"提示链id" v:"required"`
}

type DeleteChainRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type CreateChainReq struct {
	g.Meta       `path:"/chain" tags:"提示链相关接口" method:"post" summary:"创建提示链"`
	Name         string                   `json:"name" dc:"流程名称" v:"required"`
	Nodes        []map[string]interface{} `json:"nodes" dc:"节点信息" `
	Edges        []map[string]interface{} `json:"edges" dc:"边信息"`
	IsPublic     uint                     `json:"isPublic" dc:"是否公开"`
	UseKnowledge uint                     `json:"useKnowledge" dc:"是否使用知识库" ` // 是否使用知识库
}

type CreateChainRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type ModifyChainReq struct {
	g.Meta `path:"/chain" tags:"提示链相关接口" method:"put" summary:"更新提示链"`
	Id     uint                     `json:"id" dc:"提示链id" v:"required"`
	Nodes  []map[string]interface{} `json:"nodes" dc:"节点信息" `
	Edges  []map[string]interface{} `json:"edges" dc:"边信息"`
}

type ModifyChainRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type UpdateChainNameReq struct {
	g.Meta `path:"/chain/updateChainName" tags:"提示链相关接口" method:"put" summary:"更新提示链名称"`
	Id     uint   `json:"id" dc:"提示链id" v:"required"`
	Name   string `json:"name" dc:"流程名称" v:"required"`
}

type UpdateChainNameRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type ChainBotListReq struct {
	g.Meta  `path:"/chain/chainBotList" tags:"提示链相关接口" method:"post" summary:"获取chain关联的bot"`
	ChainId uint `json:"chainId" dc:"提示链id" v:"required"`
}

type ChainBotListRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.Bot `json:"data" dc:"数据体"`
}
