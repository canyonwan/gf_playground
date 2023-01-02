package frontend

import (
	"gf_playground/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TodoResBase 响应体
type TodoResBase struct {
	Title     string      `json:"title" dc:"标题"`
	Content   string      `json:"content" dc:"内容"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间"`
}

// TodoReqBase 请求体
type TodoReqBase struct {
	Title   string `json:"title" dc:"标题" v:"required#标题不能为空"`
	Content string `json:"content" dc:"内容"`
}

// TodoPageGetReq 列表分页req
type TodoPageGetReq struct {
	g.Meta `path:"frontend/todo/page" method:"get" tags:"待办事项" summary:"分页列表"`
	common.PageCommonReq
}

// TodoPageGetRes 列表分页res
type TodoPageGetRes struct {
	common.PageCommonRes
}

// TodoListGetReq 不分页 列表
type TodoListGetReq struct {
	g.Meta `path:"frontend/todo/list" method:"get" tags:"待办事项" summary:"列表"`
}
type TodoListGetRes struct {
	List interface{} `json:"list" dc:"列表"`
}

// TodoCreateReq 新增
type TodoCreateReq struct {
	g.Meta `path:"frontend/todo" method:"post" tags:"待办事项" summary:"新增"`
	TodoReqBase
}
type TodoCreateRes struct {
	Id int `json:"id" dc:"id"`
}

// TodoUpdateReq 编辑
type TodoUpdateReq struct {
	g.Meta `path:"frontend/todo" method:"put" tags:"待办事项" summary:"编辑"`
	Id     int `json:"id" dc:"id" v:"required#id不能为空"`
	TodoReqBase
}
type TodoUpdateRes struct {
	Id int `json:"id" desc:"id"`
}

// TodoDeleteReq 删除
type TodoDeleteReq struct {
	g.Meta `path:"frontend/todo/{id}" in:"path" method:"delete" tags:"待办事项" summary:"删除"`
	Id     int `json:"id" v:"required#id不能为空" dc:"Todo Id" `
}
type TodoDeleteRes struct{}
