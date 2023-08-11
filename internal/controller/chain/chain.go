package chain

import (
	"context"
	v1 "llmfarm/api/chain/v1"
	botService "llmfarm/internal/service/bot"
	service "llmfarm/internal/service/chain"
	"llmfarm/internal/service/session"
	"llmfarm/library/response"

	"github.com/gogf/gf/v2/frame/g"
)

type ChainController struct {
}

func New() *ChainController {
	return &ChainController{}
}

func (c *ChainController) Post(ctx context.Context, req *v1.CreateChainReq) (res *v1.CreateChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	chainInfo, err := service.New().Create(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", chainInfo)
	return
}

func (c *ChainController) Get(ctx context.Context, req *v1.DetailChainReq) (res *v1.DetailChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	chainInfo, err := service.New().Detail(req.Id, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", chainInfo)
	return
}

func (c *ChainController) List(ctx context.Context, req *v1.ListChainReq) (res *v1.ListChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	chainList, err := service.New().List(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", chainList)
	return
}

func (c *ChainController) Delete(ctx context.Context, req *v1.DeleteChainReq) (res *v1.DeleteChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	err = service.New().Delete(req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", nil)
	return
}

func (c *ChainController) Put(ctx context.Context, req *v1.ModifyChainReq) (res *v1.ModifyChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	err = service.New().Update(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (c *ChainController) UpdateChainName(ctx context.Context, req *v1.UpdateChainNameReq) (res *v1.ModifyChainRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	err = service.New().UpdateChainName(req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (c *ChainController) ChainBotList(ctx context.Context, req *v1.ChainBotListReq) (res *v1.ChainBotListRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	bots, err := botService.New().ChainBotList(int(session.Context.GetUser(ctx).Id), int(req.ChainId))
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", bots)
	return
}
