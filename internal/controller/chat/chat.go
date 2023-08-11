package chat

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "llmfarm/api/chat/v1"
	"llmfarm/internal/model/entity"
	service "llmfarm/internal/service/chat"
	"llmfarm/internal/service/session"
	"llmfarm/library/response"
)

type ChatController struct{}

func New() *ChatController {
	return &ChatController{}
}

func (c *ChatController) Post(ctx context.Context, req *v1.CreateChatReq) (res *v1.CreateChatRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	ress, err := service.New().Create(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

func (c *ChatController) Put(ctx context.Context, req *v1.ModifyChatReq) (res *v1.ModifyChatRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	var ress *entity.Chat
	ress, err = service.New().Modify(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

func (c *ChatController) Delete(ctx context.Context, req *v1.DeleteChatReq) (res *v1.DeleteChatRes, err error) {
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

func (c *ChatController) ClearChat(ctx context.Context, req *v1.ClearChatReq) (res *v1.ClearChatRes, err error) {
	err = service.New().ClearChat(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (c *ChatController) List(ctx context.Context, req *v1.ChatListReq) (res *v1.ChatListRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	ress, _ := service.New().List(req, session.Context.GetUser(ctx).Id)
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

//ModifyChatContentType

func (c *ChatController) ModifyChatContentType(ctx context.Context, req *v1.ModifyChatContextTypeReq) (res *v1.ModifyChatContextTypeRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	var ress *entity.Chat
	ress, err = service.New().ModifyChatContentType(req, session.Context.GetUser(ctx).Id)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}
