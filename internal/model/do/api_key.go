// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiKey is the golang structure of table apiKey for DAO operations like Where/Data.
type ApiKey struct {
	g.Meta     `orm:"table:apiKey, do:true"`
	Id         interface{} // 主键
	Created_at *gtime.Time // 创建时间
	Update_at *gtime.Time // 更新时间
	Deleted_at *gtime.Time // 删除时间
	Name       interface{} // 名称
	DxpirationTime       *gtime.Time // 过期时间
	EnableStatus  interface{} // 启用状态
	PaymentStatus interface{} // 付费状态
}
