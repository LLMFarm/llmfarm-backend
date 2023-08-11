package tokenUsageLimit

import (
	"context"
	v1 "llmfarm/api/tokenUsageLimit/v1"
	"llmfarm/library/response"
)

type TokenUsageLimitController struct{}

func New() *TokenUsageLimitController {
	return &TokenUsageLimitController{}
}

func (c *TokenUsageLimitController) List(ctx context.Context, req *v1.TokenUsageLimitListReq) (res *v1.TokenUsageLimitListRes, err error) {
	// versionType, _ := g.Cfg().Get(ctx, "versionType")
	list := []v1.TokenUsageLimit{}
	// list, err := service.New().List(ctx, versionType.String(), req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", list)
	return
}
