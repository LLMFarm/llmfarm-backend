// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BotSetFolder is the golang structure of table BotSetFolder for DAO operations like Where/Data.
type BotSetFolder struct {
	g.Meta     `orm:"table:bot_set_folder, do:true"`
	Id         interface{} // 主键
	CreateTime 	*gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	BotId       interface{} // 机器人Id
	SetFolderId interface{} // 文件夹Id

}