package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"llmfarm/internal/controller/bot"
	"llmfarm/internal/controller/chain"
	"llmfarm/internal/controller/chat"
	"llmfarm/internal/controller/message"
	"llmfarm/internal/controller/tokenUsageLimit"
	"llmfarm/internal/controller/user"
	"llmfarm/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// s.SetServerRoot("/Users/guwenshuai/Downloads")
			fileStorageType, _ := g.Cfg().Get(ctx, "fileStorageType")
			if fileStorageType.String() == "local" {
				filePath, _ := g.Cfg().Get(ctx, "localconfig.path")
				fmt.Println("设置静态路径:", filePath.String())
				s.SetServerRoot(filePath.String())
			}
			//设置最大请求体大小
			s.SetClientMaxBodySize(90 * 1024 * 1024)
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.Middleware.Ctx,
				)

				group.REST("/user", user.New())
				group.POST("/user/login", user.New().Login)
				group.POST("/user/signup", user.New().SignUp)
				group.POST("/user/signout", user.New().SignOut)
				group.POST("/user/loginByVerify", user.New().LoginWithVerifyCode)
				group.POST("/user/forget", user.New().ForgetPassword)
				group.Middleware(middleware.Middleware.Auth)

				group.REST("/chat", chat.New())
				group.GET("/chat/list", chat.New().List)
				group.POST("/chat/clear", chat.New().ClearChat)
				group.GET("/tokenUsageLimit/list", tokenUsageLimit.New().List)
				group.GET("/messages", message.New().List)
				group.PUT("/message/active", message.New().ModifyActive)
				group.POST("/message", message.New().Question)
				group.POST("/message/reAnswer", message.New().ReAnswer)
				group.POST("/message/stopAnswer", message.New().StopAnswer)
				group.POST("/message/evaluate", message.New().Evaluate)
				group.POST("/message/compare", message.New().Compare)
				group.GET("/message/getUserDailyUsage", message.New().GetUserDailyUsage)
				group.POST("/chat/contextType", chat.New().ModifyChatContentType)

				group.REST("/bot", bot.New())
				group.GET("/bot/list", bot.New().List)
				group.GET("/bot/getPublicBotList", bot.New().GetPublicBotList)
				group.GET("/bot/usedBotList", bot.New().GetUsedBotList)

				group.REST("/chain", chain.New())
				group.GET("/chain/list", chain.New().List)
				group.POST("/chain/updateChainName", chain.New().UpdateChainName)
				group.POST("/chain/chainBotList", chain.New().ChainBotList)

				group.GET("/getUserInfo", user.New().GetUserInfo)

				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
