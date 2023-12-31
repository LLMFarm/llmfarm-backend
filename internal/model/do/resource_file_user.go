// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFile is the golang structure of table userFile for DAO operations like Where/Data.
type UserFile struct {
	g.Meta       `orm:"table:userFile, do:true"`
	Id           interface{} //
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	DeleteTime   *gtime.Time // 删除时间
	Name         interface{} // 文件名
	UserId interface{} //用户Id
	FileId         interface{} // 文件资源ID
}
