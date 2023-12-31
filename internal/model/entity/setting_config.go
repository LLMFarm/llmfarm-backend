// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// SettingConfig is the golang structure for table settingConfig.
type SettingConfig struct {
	Id         uint        `json:"id" dc:"ID"         ` // ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	Name        string     `json:"name" dc:"名称"` // 名称
	Desc        string     `json:"desc" dc:"描述"` // 描述
	Unique      string     `json:"unique" dc:"唯一标识"` // 唯一标识
	Configs     [] map[string]interface{}     `json:"configs" dc:"配置"` // 配置
}

func (df *SettingConfig) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *SettingConfig) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}