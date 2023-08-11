package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type TokenUsageLimitDetailReq struct {
	g.Meta  `path:"/TokenUsageLimit/detail" tags:"token使用限制相关接口" method:"get" summary:"查询令牌使用限制详情"`
	ModelId uint `v:"required"`
}

type TokenUsageLimitDetailRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.TokenUsageLimit `json:"data" dc:"数据体"`
}

// TokenUsageLimitCreateReq
type TokenUsageLimitCreateReq struct {
	g.Meta      `path:"/tokenUsageLimit/create" tags:"token使用限制相关接口" method:"post" summary:"创建token使用限制"`
	ModelId     uint `json:"modelId" dc:"模型id" v:"required"`
	DailyTokens uint `json:"dailyTokens" dc:"使用限制" v:"required"`
}

type TokenUsageLimitCreateRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

// TokenUsageLimitListReq
type TokenUsageLimitListReq struct {
	g.Meta `path:"/tokenUsageLimit/list" tags:"token使用限制相关接口" method:"get" summary:"查询token使用限制列表"`
}

type TokenUsageLimitListRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []TokenUsageLimit `json:"data" dc:"数据体"`
}

type TokenUsageLimit struct {
	Title       string `json:"title" dc:"标题" `
	DailyTokens uint   `json:"dailyTokens" dc:"使用限制" `
	UsageToken  uint   `json:"usageToken" dc:"已使用令牌数" `
}
