// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)
// DataSet is the golang structure for table chat.
type DataSet struct {
	Id         uint        `json:"id" dc:"ID"         ` // ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	SetName       string        `json:"setName" dc:"数据集名称"      ` // 数据集名称
	UserId uint        `json:"userId" dc:"用户ID"      ` // 用户ID
	EmbeddingModel string        `json:"embeddingModel" dc:"嵌入模型"      ` // 嵌入模型
	VectorBase string        `json:"vectorBase" dc:"向量库"      ` // 向量库
	Params map[string]interface{}        `json:"params" dc:"参数"      ` // 参数
	UserFileId uint        `json:"userFileId" dc:"用户文件ID"      ` // 用户文件ID
	ParseStatus string        `json:"parseStatus" dc:"解析状态"      ` // 解析状态
	ParseProgress int        `json:"parseProgress" dc:"解析进度"      ` // 解析进度
}

func (df *DataSet) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *DataSet) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}