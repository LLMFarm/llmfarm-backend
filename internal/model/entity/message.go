// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id          uint         `json:"id" dc:"消息ID"         ` // 消息ID
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime" dc:"更新时间"  ` // 更新时间
	DeleteTime  *gtime.Time `json:"deleteTime" dc:"删除时间"  ` // 删除时间
	UserId      uint         `json:"userId" dc:"用户ID"      ` // 用户ID
	Type        string      `json:"type" dc:"消息发送者类型"        ` // 消息发送者类型
	Content     string      `json:"content" dc:"消息内容"    ` // 消息内容
	ContentType string      `json:"contentType" dc:"消息内容类型" ` // 消息内容类型
	ParentId    uint         `json:"parentId" dc:"父消息ID"   ` // 父消息ID
	Mark        string      `json:"mark" dc:"标记"       ` // 标记
	Feedback    string      `json:"feedback" dc:"反馈信息"    ` // 反馈信息
	Active      int         `json:"active" dc:"是否有效"      ` // 是否有效
	ChatId      uint        `json:"chatId"     ` // 所属对话ID
	BotId       uint 	    `json:"botId" dc:"机器人ID" ` //机器人ID
}

func (df *Message) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Message) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}