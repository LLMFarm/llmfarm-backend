// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Folder is the golang structure of table folder for DAO operations like Where/Data.
type Folder struct {
	g.Meta     `orm:"table:folder, do:true"`
	Id         interface{} // 模型ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	Name       interface{} // 资源名称
	Url       interface{} // 资源链接
	BotId  	  interface{} // 机器人ID
	UserId     interface{} // 用户ID
	parseStatus  interface{} // 解析状态
}
