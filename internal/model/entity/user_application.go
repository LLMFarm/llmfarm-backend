// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserApplication is the golang structure for table UserApplication.
type UserApplication struct {
	Id          uint        `json:"id"  dc:"ID"          ` // ID
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"   ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime" dc:"更新时间"   ` // 更新时间
	DeleteTime  *gtime.Time `json:"deleteTime" dc:"删除时间"` // 删除时间
	UserName  string `json:"userName" dc:"姓名"` // 姓名
	Phone string `json:"phone" dc:"手机号"` // 手机号
	EnterpriseName string `json:"enterpriseName" dc:"企业名称"` // 企业名称
	Post string `json:"post" dc:"职位"` // 职位
	Purpose string `json:"purpose" dc:"用途"` // 用途
	RequirementScenario string `json:"requirementScenario" dc:"需求场景"` // 需求场景
	Status string `json:"status" dc:"申请状态 申请成功 申请中 "` // 状态
	UserId uint `json:"userId" dc:"用户id"` // 用户id
}

func (df *UserApplication) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *UserApplication) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}