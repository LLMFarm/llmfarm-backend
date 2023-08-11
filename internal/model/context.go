package model

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/net/ghttp"
)

const ContextKey = "ChatContext"

type ContextUser struct {
	Id              uint        `json:"id"       `                     //
	Name            string      `json:"name"       `                   //
	Phone           string      `json:"phone"      `                   //
	Email           string      `json:"email"      `                   //
	OpenId          string      `json:"openId"     `                   //
	Is_admin        int         `json:"is_admin"   `                   // 是否是管理员
	Is_membership   int         `json:"is_membership"   dc:"是否是会员"`    // 是否是会员
	Member_deadline *gtime.Time `json:"member_deadline"   dc:"会员到期时间"` // 会员到期时间
}

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}
