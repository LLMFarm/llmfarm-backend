package middleware

import (
	"context"
	"llmfarm/internal/model"
	"llmfarm/internal/service/session"
	"llmfarm/library/response"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: g.RequestFromCtx(r.Context()).Session,
	}
	session.Context.Init(r.Context(), customCtx)
	if user := session.Session.GetUser(g.RequestFromCtx(r.Context()).Context()); user != nil {

		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Name:     user.Name,
			Phone:    user.Phone,
			Email:    user.Email,
			Is_admin: user.Is_admin,
			// OpenId: user.OpenId,
		}
	}
	// 执行下一步请求逻辑
	g.RequestFromCtx(r.Context()).Middleware.Next()
}
func (s *middlewareService) IsSignedIn(ctx context.Context) bool {
	return session.Context.GetUser(ctx) != nil
}

// GetUserApplication
func (s *middlewareService) Auth(r *ghttp.Request) {
	if s.IsSignedIn(r.Context()) {
		//通过校验后刷新过期时间
		cc, _ := g.RequestFromCtx(r.Context()).Request.Cookie("gfsessionid")
		if cc != nil {
			cookie := http.Cookie{
				Name:     "gfsessionid",
				Value:    cc.Value,
				Expires:  time.Now().Add(32 * time.Hour),
				Path:     "/",
				Domain:   "",
				Secure:   false,
				HttpOnly: false,
			}
			http.SetCookie(g.RequestFromCtx(r.Context()).Response.Writer, &cookie)
		}
		g.RequestFromCtx(r.Context()).Middleware.Next()
	} else {
		response.JsonFailExit(r.Context(), "需要登录")
	}
}
