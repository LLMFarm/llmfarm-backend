package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateChatReq struct {
	g.Meta      `path:"/chat" tags:"会话相关接口" method:"post" summary:"创建一个会话"`
	Name        string `json:"name" dc:"会话名称" v:"required"`
	Type        string `json:"type" dc:"会话类型" v:"required"`
	ContextType string `json:"contextType" dc:"会话上下文类型"`
	IsTemp      int    `json:"isTemp" dc:"是否是临时会话 0否 1是"`
}
type CreateChatRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Chat `json:"data" dc:"数据体"`
}

type ChatListReq struct {
	g.Meta  `path:"/chat/list" tags:"会话相关接口" method:"get" summary:"查询会话列表"`
	Keyword string `json:"keyword" dc:"会话名称关键字搜索"`
	Type    string `json:"type" dc:"会话类型（CHAT/QA/EXCEL）默认CHAT" df:"CHAT"`
}

type ChatListRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.Chat `json:"data" dc:"数据体"`
}

type ModifyChatReq struct {
	g.Meta `path:"/chat" tags:"会话相关接口" method:"put" summary:"修改会话"`
	Name   string `json:"name" dc:"会话名称" v:"required"`
	ChatId uint   `json:"chatId" dc:"会话id" v:"required"`
}

type ModifyChatContextTypeReq struct {
	g.Meta      `path:"/chat/contextType" tags:"会话相关接口" method:"put" summary:"修改会话上下文类型"`
	ContextType string `json:"contextType" dc:"会话类型" v:"required"`
	ChatId      uint   `json:"chatId" dc:"会话id" v:"required"`
}

type ModifyChatContextTypeRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Chat `json:"data" dc:"数据体"`
}

type ModifyChatRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Chat `json:"data" dc:"数据体"`
}

type DeleteChatReq struct {
	g.Meta `path:"/chat" tags:"会话相关接口" method:"delete" summary:"删除会话"`
	ChatId uint `json:"chatId" dc:"会话id" v:"required"`
}

type DeleteChatRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type ClearChatReq struct {
	g.Meta `path:"/chat/clear" tags:"会话相关接口" method:"post" summary:"清除会话"`
}

type ClearChatRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}
