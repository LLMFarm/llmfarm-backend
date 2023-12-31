// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// Bot is the golang structure for table bot.
type Bot struct {
	Id         uint        `json:"id" dc:"机器人ID"         ` // 聊天ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	BotName        string 	   `json:"botName" dc:"机器人名称"` // 机器人名称
	ChainId uint     `json:"chainId" dc:"提示链id"` // 提示链id
	UserType  string     `json:"userType" dc:"用户类型"` // 用户类型
	UserId    uint      	`json:"userId" dc:"用户Id"` // 用户Id
	Icon string `json:"icon" dc:"图标"`
	IsDefault int `json:"isDefault" dc:"是否默认选中 0是false 1是true"`
}

func (df *Bot) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Bot) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}