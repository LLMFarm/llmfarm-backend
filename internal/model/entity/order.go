// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id                   uint        `json:"id" dc:"订单ID"                   ` // 订单ID
	CreateTime           *gtime.Time `json:"createTime"  dc:"创建时间"          ` // 创建时间
	UpdateTime           *gtime.Time `json:"updateTime" dc:"更新时间"           ` // 更新时间
	DeleteTime           *gtime.Time `json:"deleteTime" dc:"删除时间"           ` // 删除时间
	UserId               uint        `json:"userId" dc:"用户ID"              ` // 用户ID
	OrderNumber          string      `json:"orderNumber" dc:"订单号"         ` // 订单号
	ProductId            uint        `json:"productId"  dc:"商品ID"          ` // 商品ID
	ProductName          string      `json:"productName"  dc:"商品名称"        ` // 商品名称
	ProductDescription   string      `json:"productDescription" dc:"商品描述"  ` // 商品描述
	Amount               float64     `json:"amount"   dc:"订单总额"            ` // 商品数量
	ProductPrice         float64     `json:"productPrice"  dc:"商品单价"       ` // 商品单价
	ProductOriginalPrice float64     `json:"productOriginalPrice" dc:"商品原价" ` // 商品原价
	Status               string      `json:"status" dc:"订单状态"              ` // 订单状态
	PaymentTime          *gtime.Time `json:"paymentTime" dc:"支付时间"         ` // 支付时间
	StartTime            *gtime.Time `json:"startTime" dc:"订单开始时间"` 
	EndTime              *gtime.Time `json:"endTime" dc:"订单结束时间"`   
	CodeUrl				 string      `json:"codeUrl" dc:"付款二维码"`
}

func (df *Order) DefaultCreateTime() {
	df.CreateTime = gtime.Now()
}

// DefaultCreateAt changes the default createAt field
func (df *Order) DefaultUpdateTime() {
	df.UpdateTime = gtime.Now()
}