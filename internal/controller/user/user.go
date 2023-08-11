package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"

	v1 "llmfarm/api/user/v1"
	"llmfarm/internal/model/entity"
	"llmfarm/internal/service/session"
	service "llmfarm/internal/service/user"
	"llmfarm/library/response"
)

type UserController struct{}

func New() *UserController {
	return &UserController{}
}

func (c *UserController) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var ress *v1.UserLoginResult
	ress, err = service.New().Login(ctx, req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

func (c *UserController) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	_, err = service.New().SignUp(ctx, req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (c *UserController) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	session.Session.RemoveUser(ctx)
	// service.New().SignOut(ctx, req)
	return
}

func (c *UserController) LoginWithVerifyCode(ctx context.Context, req *v1.UserLoginWithVerifyCodeReq) (res *v1.UserLoginWithVerifyCodeRes, err error) {
	var ress *v1.UserLoginResult
	ress, err = service.New().LoginWithVerifyCode(ctx, req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success", ress)
	return
}

func (c *UserController) ForgetPassword(ctx context.Context, req *v1.UserForgetPasswordReq) (res *v1.UserForgetPasswordRes, err error) {
	_, err = service.New().ForgetPassword(ctx, req)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	response.JsonSuccessExit(ctx, "success")
	return
}

func (c *UserController) Get(ctx context.Context, req *v1.FindUserReq) (res *v1.FindUserRes, err error) {
	if err := g.RequestFromCtx(ctx).Parse(&req); err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	// ress, _ := service.New().Detail(req)

	var ctxUser *entity.User
	gconv.Struct(g.Map{
		"Id":            1,
		"Name":          "llmfarm",
		"Phone":         "15910771917",
		"Email":         "",
		"OpenId":        "",
		"Is_admin":      1,
		"Is_membership": 0,
	}, &ctxUser)
	session.Session.SetUser(ctx, ctxUser)

	user := session.Context.GetUser(ctx)
	fmt.Println(user)
	if user == nil || user.Id == 0 {
		response.JsonFailExit(ctx, "未登录")
	}
	response.JsonSuccessExit(ctx, "success", user)
	return
}

// GetUserInfoReq
func (c *UserController) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	//session.Context.GetUser(ctx).Id
	ress, _ := service.New().GetUserById(session.Context.GetUser(ctx).Id)
	response.JsonSuccessExit(ctx, "success", ress)
	return
}
