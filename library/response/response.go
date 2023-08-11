package response

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	// 处理国际化 message
	//g.I18n().WithLanguage(r.Context(), r.Header.Get("Content-Language"))
	message = g.I18n().T(r.Context(), message)
	fmt.Println("message", message)
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(ctx context.Context, err int, msg string, data ...interface{}) {
	Json(g.RequestFromCtx(ctx), err, msg, data...)
	g.RequestFromCtx(ctx).Exit()
}

func JsonSuccessExit(ctx context.Context, msg string, data ...interface{}) {
	JsonExit(ctx, 0, msg, data...)
}

func JsonFailExit(ctx context.Context, msg string, data ...interface{}) {
	JsonExit(ctx, 1, msg, data...)
}
