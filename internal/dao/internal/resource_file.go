// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResourceFileDao is the data access object for table file.
type ResourceFileDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ResourceFileColumns // columns contains all the column names of Table for convenient usage.
}

// ResourceFileColumns defines and stores column names for table file.
type ResourceFileColumns struct {
	Id           string //
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	DeleteTime   string // 删除时间
	Url          string // 文件地址
	Name         string // 文件名
	ParseStatus  string // 解析状态 1：Parsing  2:Success 3:Failure
	Md5          string // 文件的md5值
	CreateUserId string //
	Size         string // 文件大小，单位kb
}

// resourceFileColumns holds the columns for table file.
var resourceFileColumns = ResourceFileColumns{
	Id:           "id",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	DeleteTime:   "deleteTime",
	Url:          "url",
	Name:         "name",
	ParseStatus:  "parseStatus",
	Md5:          "md5",
	CreateUserId: "createUserId",
	Size:         "size",
}

// NewResourceFileDao creates and returns a new DAO object for table data access.
func NewResourceFileDao() *ResourceFileDao {
	return &ResourceFileDao{
		group:   "default",
		table:   "file",
		columns: resourceFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ResourceFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ResourceFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ResourceFileDao) Columns() ResourceFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ResourceFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ResourceFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ResourceFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
