// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApproveRecord is the golang structure of table approveRecord for DAO operations like Where/Data.
type ApproveRecord struct {
	g.Meta     `orm:"table:approveRecord, do:true"`
	Id         interface{} // 记录ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	UserId     interface{} // 用户ID
	TemplateId interface{} // 关联的模版id
	Approver   interface{} // 审批人id
	Result     interface{} // 审批结果 1:通过 2:驳回 3:已上架 4:已下架
	Content    interface{} // 审批备注
	IsRead     interface{} // 是否已读 0未读 1已读
}
