// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// Model is the golang structure for table chat.
type Model struct {
	Id         uint        `json:"id" dc:"模型ID"         ` // 模型ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	Name       string      `json:"name" dc:"模型名称"      ` // 模型名称
	Desc       string      `json:"type" dc:"模型描述"      ` // 模型描述
	SortOrder  uint        `json:"sortOrder" dc:"排序"    ` // 排序
	MaxTokenLimit int      `json:"maxTokenLimit" dc:"最大token限制"`
	Tag        string      `json:"tag" dc:"模型唯一标识"`
	Factor     float32      `json:"factor" dc:"因子"`
}

func (df *Model) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Model) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}