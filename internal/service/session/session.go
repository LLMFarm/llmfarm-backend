package session

import (
	"context"
	"fmt"
	"llmfarm/internal/model/entity"
)

// Session管理服务
var Session = sessionService{}

type sessionService struct{}

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "SESSIONID"
)

// 设置用户Session.
func (s *sessionService) SetUser(ctx context.Context, user *entity.User) error {
	fmt.Println("user", user)
	fmt.Println(Context.Get(ctx), "===")
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sessionService) GetUser(ctx context.Context) *entity.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v, err := customCtx.Session.Get(sessionKeyUser, ""); !v.IsNil() && err == nil {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// 删除用户Session。
func (s *sessionService) RemoveUser(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
