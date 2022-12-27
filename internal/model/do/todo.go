// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Todo is the golang structure of table todo for DAO operations like Where/Data.
type Todo struct {
	g.Meta    `orm:"table:todo, do:true"`
	Id        interface{} // id
	Title     interface{} // 标题
	Content   interface{} // 详细内容
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
}
