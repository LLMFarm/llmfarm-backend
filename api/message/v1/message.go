package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type EvaluateReq struct {
	g.Meta    `path:"/message/evaluate" tags:"消息相关接口" method:"post" summary:"评价回答"`
	Mark      string `json:"mark" dc:"标记" v:"required"`
	Feedback  string `json:"feedback" dc:"反馈" `
	MessageId string `json:"messageId" dc:"问答id" v:"required"`
}

type EvaluateRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type CompareReq struct {
	g.Meta          `path:"/message/compare" tags:"消息相关接口" method:"post" summary:"对比回答"`
	SourceMessageId uint   `json:"messageId" dc:"源问答id" v:"required"`
	ParentId        uint   `json:"parentId" dc:"父id" v:"required"`
	TargetMessageId uint   `json:"targetMessageId" dc:"目标问答id" v:"required"`
	Result          string `json:"result" dc:"对比结果" v:"required"`
}

type CompareRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.CompareRecord `json:"data" dc:"数据体"`
}

type Message struct {
	*entity.Message
	BotName string `json:"botName" `
	Icon    string `json:"icon"`
}

type MessageListReq struct {
	g.Meta `path:"/messages" tags:"消息相关接口" method:"get" summary:"获取消息列表"`
	ChatId uint `v:"required" json:"chatId" dc:"会话id"`
}
type MessageListRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data []*entity.Message `json:"data" dc:"数据体"`
}

type UpdateMessageReq struct {
	g.Meta `path:"/message/active" tags:"消息相关接口" method:"put" summary:"修改消息active"`
	Id     uint `v:"required" json:"chatId" dc:"会话id"`
}
type UpdateMessageRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.Message `json:"data" dc:"数据体"`
}

type ConversationReq struct {
	g.Meta      `path:"/message" tags:"消息相关接口" method:"post" summary:"提问"`
	ChatId      uint   `json:"chatId" dc:"会话id"`
	ParentId    uint   `json:"parentId" dc:"父消息Id"`
	Content     string `v:"required" json:"content" dc:"消息内容"`
	BotId       uint   `v:"required" json:"botId" dc:"机器人id"`
	ContextType string `json:"contextType" dc:"上下文类型"`
	IsTemp      int    `json:"isTemp" dc:"是否是临时会话 0否 1是"`
}
type ConversationRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.Message `json:"data" dc:"数据体"`
}

type ReconsiderAnswerReq struct {
	g.Meta      `path:"/message/reAnswer" tags:"消息相关接口" method:"post" summary:"重新回答"`
	ParentId    uint   `v:"required" json:"parentId" dc:"父消息Id"`
	ChatId      uint   `v:"required" json:"chatId" dc:"会话id"`
	BotId       uint   `v:"required" json:"botId" dc:"机器人id"`
	ContextType string `json:"contextType" dc:"上下文类型"`
}

type UserSignOutReq struct {
	g.Meta  `path:"/message/getUserDailyUsage" method:"get" tags:"获取用户已用token" summary:"获取用户已用token"`
	ModelId uint `v:"required" json:"botId" dc:"模型id"`
}

type UserSignOutRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
}

type ReconsiderAnswerRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.Message `json:"data" dc:"数据体"`
}

type StopAnswerReq struct {
	g.Meta `path:"/message/stopAnswer" tags:"消息相关接口" method:"post" summary:"停止回答"`
	ChatId uint ` json:"chatId" dc:"会话id"`
}
type StopAnswerRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.Message `json:"data" dc:"数据体"`
}
