package bot

import (
	"context"
	v1 "llmfarm/api/bot/v1"
	service "llmfarm/internal/service/bot"
	"llmfarm/internal/service/session"
	"llmfarm/library/response"

	"github.com/gogf/gf/v2/frame/g"
)

type BotController struct {
}

func New() *BotController {
	return &BotController{}
}

// 获取我能用的bot
func (c *BotController) List(ctx context.Context, req *v1.ListBotReq) (res *v1.ListBotRes, err error) {
	botInfo, err := service.New().List(req.Word, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", botInfo)
	return
}

// GetSystemBotList
func (c *BotController) GetPublicBotList(ctx context.Context, req *v1.GetPublicBotListReq) (res *v1.GetPublicBotListRes, err error) {
	botList, pageCount, err := service.New().GetPublicBotList(req.Word, session.Context.GetUser(ctx).Id, req.Limit, req.PageNumber, req.PromptType, req.ScenarioId)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	var botData = new(v1.BotListData)
	botData.BotList = botList
	botData.PageCount = pageCount

	response.JsonSuccessExit(ctx, "success", botData)
	return
}

func (c *BotController) Get(ctx context.Context, req *v1.DetailBotReq) (res *v1.DetailBotRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	botInfo, err := service.New().Detail(req.BotId, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", botInfo)
	return
}

func (c *BotController) Post(ctx context.Context, req *v1.CreateBotReq) (res *v1.CreateBotRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	botInfo, err := service.New().Create(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", botInfo)
	return
}

func (c *BotController) Put(ctx context.Context, req *v1.ModifyBotReq) (res *v1.ModifyBotRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	botInfo, err := service.New().Modify(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", botInfo)
	return
}

func (c *BotController) Delete(ctx context.Context, req *v1.DeleteBotReq) (res *v1.DeleteBotRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	err = service.New().Delete(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

// 获取最近使用过的机器人
func (c *BotController) GetUsedBotList(ctx context.Context, req *v1.GetUsedBotListReq) (res *v1.GetUsedBotListRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	list, err := service.New().UsedBots(session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", list)
	return
}
