package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// getUserInfo
type GetUserInfoReq struct {
	g.Meta `path:"/user/getUserInfo" tags:"用户相关接口" method:"get" summary:"获取用户信息"`
}

type GetUserInfoRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.User `json:"data" dc:"数据体"`
}

type FindUserReq struct {
	g.Meta `path:"/user/get" tags:"用户相关接口" method:"get" summary:"根据手机号或邮箱查询用户"`
	Phone  string `json:"phone" dc:"手机号码"`
	Email  string `json:"email" dc:"邮箱"`
}
type FindUserRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data entity.User `json:"data" dc:"数据体"`
}

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户相关接口" method:"post" summary:"用户登录"`
	Username string `json:"username" dc:"手机号或邮箱" v:"required"`
	Password string `json:"password" dc:"用户密码" v:"required"`
}

type UserLoginResult struct {
	UserId   uint   `json:"userId" dc:"用户id"`
	Name     string `json:"name" dc:"用户名"`
	Phone    string `json:"phone" dc:"用户手机号"`
	Email    string `json:"email" dc:"邮件地址"`
	Is_admin int    `json:"is_admin" dc:"是否是管理员"`
}

// LoginGoogle
type LoginGoogleReq struct {
	g.Meta  `path:"/user/loginGoogle" tags:"用户相关接口" method:"post" summary:"用户登录"`
	IdToken string `json:"idToken" dc:"idToken" v:"required"`
}

type LoginGoogleRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data UserLoginResult `json:"data" dc:"数据体"`
}

type UserLoginRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data UserLoginResult `json:"data" dc:"数据体"`
}

type UserSignUpReq struct {
	g.Meta         `path:"/user/signup" tags:"用户相关接口" method:"post" summary:"用户注册"`
	Phone          string `json:"phone" v:"required"   dc:"手机号"`
	Code           string `json:"code"  v:"required"   dc:"验证码"`
	Password       string `json:"password" v:"required" dc:"用户密码"`
	InvitationCode string `json:"invitationCode" dc:"邀请码" v:"required" `
}

type UserSignUpRes struct {
	g.Meta `mime:"application/json" example:"string"`
	UserId uint   `json:"userId" dc:"用户id"`
	Name   string `json:"name" dc:"用户名"`
	Phone  string `json:"phone" dc:"用户手机号"`
	Email  string `json:"email" dc:"邮件地址"`
	api.Common
}

type UserSignOutReq struct {
	g.Meta `path:"/user/signout" method:"post" tags:"用户相关接口" summary:"用户登出"`
}
type UserSignOutRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
}

type SendUserVerificationCodeReq struct {
	g.Meta `path:"/user/send" tags:"用户相关接口" method:"post" summary:"发送验证码"`
	Phone  string `json:"phone"  v:"required"  dc:"手机号"`
}

type SendUserVerificationCodeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
}

type UserLoginWithVerifyCodeReq struct {
	g.Meta `path:"/user/loginByVerify" tags:"用户相关接口" method:"post" summary:"用户使用验证码登录"`
	Phone  string `json:"phone" v:"required"   dc:"手机号"`
	Code   string `json:"code"  v:"required"   dc:"验证码"`
}

type UserLoginWithVerifyCodeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	api.Common
	Data UserLoginResult `json:"data" dc:"数据体"`
}

type UserForgetPasswordReq struct {
	g.Meta   `path:"/user/forget" tags:"用户相关接口" method:"post" summary:"用户忘记密码"`
	Phone    string `json:"phone" v:"required"   dc:"手机号"`
	Code     string `json:"code"  v:"required"   dc:"验证码"`
	Password string `json:"password" v:"required" dc:"密码"`
}

type UserForgetPasswordRes struct {
	g.Meta `mime:"application/json" example:"string"`
	UserId uint   `json:"userId" dc:"用户id"`
	Name   string `json:"name" dc:"用户名"`
	Phone  string `json:"phone" dc:"用户手机号"`
	Email  string `json:"email" dc:"邮件地址"`
	api.Common
}
