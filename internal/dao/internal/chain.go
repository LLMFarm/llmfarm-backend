// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChainDao is the data access object for table chain.
type ChainDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns ChainColumns // columns contains all the column names of Table for convenient usage.
}

// ChainColumns defines and stores column names for table chain.
type ChainColumns struct {
	Id                  string // 流程ID
	CreateTime          string // 创建时间
	UpdateTime          string // 更新时间
	DeleteTime          string // 删除时间
	UserId              string // 用户ID
	ChainName           string // 流程名称
	Nodes               string // 节点数据
	Edges               string // 边缘数据
	UserType            string // 用户类型
	UseKnowledge        string // 是否使用知识库
	ChainLevel          string // 层级
	FatherId            string // 父ID，记录所属模版
	FatherTemplateId    string // 记录所属具体某个vesion的模版
	IsUpload            string // 0 没上传 1 已上传
	ChainTemplateId     string // 上传后的模版ID
	IsInMarket          string // 是否已上架 0:未上架 1:已上架
	ChainTemplateInfoId string // 上架后的模版ID
}

// chainColumns holds the columns for table chain.
var chainColumns = ChainColumns{
	Id:                  "id",
	CreateTime:          "createTime",
	UpdateTime:          "updateTime",
	DeleteTime:          "deleteTime",
	UserId:              "userId",
	ChainName:           "chainName",
	Nodes:               "nodes",
	Edges:               "edges",
	UserType:            "userType",
	UseKnowledge:        "useKnowledge",
	ChainLevel:          "chainLevel",
	FatherId:            "fatherId",
	FatherTemplateId:    "fatherTemplateId",
	IsUpload:            "isUpload",
	ChainTemplateId:     "chainTemplateId",
	IsInMarket:          "isInMarket",
	ChainTemplateInfoId: "chainTemplateInfoId",
}

// NewChainDao creates and returns a new DAO object for table data access.
func NewChainDao() *ChainDao {
	return &ChainDao{
		group:   "default",
		table:   "chain",
		columns: chainColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChainDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChainDao) Columns() ChainColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChainDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
