// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageModelUsageTokens is the golang structure of table messageModelUsageTokens for DAO operations like Where/Data.
type MessageModelUsageTokens struct {
	g.Meta      `orm:"table:message, do:true"`
	Id          interface{} // 消息ID
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	DeleteTime  *gtime.Time // 删除时间
	MessageId   interface{} // 所属消息id
	ModelId 	interface{} // 所属模型id
	OriginalToken interface{} //当前messsage的token计算总量
	UsageTotalTokens interface{} //当前messsage的token使用总量
}
