package session

import (
	"context"
	"llmfarm/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type contextService struct{}

var Context = contextService{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(ctx context.Context, customCtx *model.Context) {
	g.RequestFromCtx(ctx).SetCtxVar(model.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 获取简略信息
func (s *contextService) GetUser(ctx context.Context) *model.ContextUser {
	value := s.Get(ctx)
	if value == nil {
		return nil
	}

	return value.User
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
