// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta      `orm:"table:goods, do:true"`
	Id          interface{} // 商品表id
	Name        interface{} // 商品名称
	Description interface{} // 商品描述
	CreateAt    interface{} //
	UpdateAt    interface{} //
}