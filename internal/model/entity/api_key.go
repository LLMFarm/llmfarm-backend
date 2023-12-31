// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// ApiKey is the golang structure for table chat.
type ApiKey struct {
	Id         uint        `json:"id" dc:"ID"         ` // ID
	Created_at *gtime.Time `json:"created_at" dc:"创建时间" ` // 创建时间
	Update_at *gtime.Time `json:"update_at" dc:"更新时间" ` // 更新时间
	Deleted_at *gtime.Time `json:"delete_at" dc:"删除时间" ` // 删除时间
	Name       string      `json:"name" dc:"名称"      ` // 名称
	ExpirationTime       *gtime.Time `json:"expiration_time" dc:"过期时间" ` // 过期时间
	EnableStatus  uint        `json:"enable_status" dc:"启用状态"    ` // 启用状态
	PaymentStatus uint        `json:"payment_status" dc:"付费状态"    ` // 付费状态

}

func (df *ApiKey) DefaultUpdateTime() {
	df.Update_at = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *ApiKey) DefaultCreateTime() {
	df.Created_at = gtime.Now()
}