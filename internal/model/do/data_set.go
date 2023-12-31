// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DataSet is the golang structure of table DataSet for DAO operations like Where/Data.
type DataSet struct {
	g.Meta     `orm:"table:data_set, do:true"`
	Id         interface{} // 主键
	CreateTime 	*gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	SetName       interface{} // 名称
	VectorBase interface{} // 向量库
	UserId interface{} // 用户ID
	Params interface{} // 参数
	UserFileId interface{} // 用户文件ID
	ParseStatus interface{} // 解析状态
	ParseProgress interface{} // 解析进度
}
