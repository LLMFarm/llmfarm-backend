// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VerificationCode is the golang structure of table verificationCode for DAO operations like Where/Data.
type VerificationCode struct {
	g.Meta     `orm:"table:verificationCode, do:true"`
	Id         interface{} // 验证码ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 更新时间
	Phone      interface{} // 手机号码
	Code       interface{} // 验证码
	ExpireTime *gtime.Time // 过期时间
	IsVerified interface{} // 是否已验证
}
