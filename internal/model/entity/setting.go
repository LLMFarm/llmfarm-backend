// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// Setting is the golang structure for table setting.
type Setting struct {
	Id         uint        `json:"id" dc:"ID"         ` // ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	ServiceType        string     `json:"serviceType" dc:"服务类型"` // 服务类型
	Value       map[string]interface{}     `json:"value" dc:"值"` // 值
	ChainId uint      	` json:"chainId" dc:"链id"` // 链id
	BotId   uint  ` json:"botId" dc:"机器人id" ` //机器人id
}

func (df *Setting) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Setting) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}