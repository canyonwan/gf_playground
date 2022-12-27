package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TodoResBase 响应体
type TodoResBase struct {
	Title     string      `json:"title" desc:"标题"`
	Content   string      `json:"content" desc:"内容"`
	CreatedAt *gtime.Time `json:"created_at" desc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" desc:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" desc:"删除时间"`
}

// TodoReqBase 请求体
type TodoReqBase struct {
	Title   string `json:"title" desc:"标题" v:"required#标题不能为空"`
	Content string `json:"content" desc:"内容" v:"required#内容不能为空"`
}

type TodoGetListCommonRes struct {
	List  interface{} `json:"list" desc:"列表"`
	Page  int         `json:"page" desc:"页码"`
	Size  int         `json:"size" desc:"页数"`
	Total int         `json:"total" desc:"总数"`
}

type TodoGetReq struct {
	g.Meta `path:"admin/todo" method:"get" tags:"待办事项" summary:"待办列表" desc:"获取待办列表"`
}
type TodoGetRes struct {
	Id int `json:"id"`
	TodoResBase
}

// TodoCreateReq 新增
type TodoCreateReq struct {
	g.Meta `path:"admin/todo" method:"post" tags:"待办事项" summary:"summary新增待办" desc:"doc新增待办"`
	TodoReqBase
}
type TodoCreateRes struct {
	Id int `json:"id" desc:"id"`
}

// TodoUpdateReq 编辑
type TodoUpdateReq struct {
	g.Meta `path:"admin/todo" method:"put" tags:"待办事项" summary:"summary编辑待办" desc:"doc编辑待办"`
	Id     int `json:"id" desc:"id" v:"required#id不能为空"`
	TodoReqBase
}
type TodoUpdateRes struct {
	Id int `json:"id" desc:"id"`
}

// TodoDeleteReq 删除
type TodoDeleteReq struct {
	g.Meta `path:"admin/todo/:id" method:"delete" tags:"待办事项" summary:"summary删除待办" desc:"desc删除待办"`
	Id     int `json:"id" desc:"id" v:"required#id不能为空"`
}
type TodoDeleteRes struct{}
