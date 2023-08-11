// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Chain is the golang structure for table chain.
type Chain struct {
	Id         uint        `json:"id" dc:"流程ID"         ` // 聊天ID
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" dc:"更新时间" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" dc:"删除时间" ` // 删除时间
	UserId uint `json:"userId" dc:"用户ID" ` // 用户ID
	ChainName string `json:"chainName" dc:"流程名称" ` // 流程名称
	Nodes  []map[string]interface{} `json:"nodes" dc:"节点" ` // 节点
	Edges  []map[string]interface{} `json:"edges" dc:"边" ` // 边
	UserType  string     `json:"userType" dc:"用户类型"` // 用户类型
	ChainLevel  int     `json:"chainLevel" dc:"层级"` 
	UseKnowledge uint `json:"useKnowledge" dc:"是否使用知识库" ` // 是否使用知识库
	FatherId int `json:"fatherId" dc:"父id，记录所属市场模版"`
	FatherTemplateId int `json:"fatherTemplateId" dc:"记录所属具体某个version的模版"`
	IsUpload int `json:"isUpload" dc:"是否已上传"`
	ChainTemplateId int `json:"chainTemplateId" dc:"上传后对应的模版id"`
	IsInMarket int `json:"isInMarket" dc:"是否已上架"`
	ChainTemplateInfoId int `json:"chainTemplateInfoId" dc:"上架后对应的市场模版id"`
}

func (df *Chain) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Chain) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}