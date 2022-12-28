package model

import "github.com/gogf/gf/v2/os/gtime"

type TodoItemBase struct {
	Title   string
	Content string
}
type TodoOutputBase struct {
	Title     string
	Content   string
	CreatedAt gtime.Time
	UpdatedAt gtime.Time
	DeletedAt gtime.Time
}

// TodoGetInput 获取
type TodoGetInput struct{}
type TodoGetOutput struct {
	Id int
	TodoOutputBase
}

// TodoCreateInput 新增
type TodoCreateInput struct {
	TodoOutputBase
}
type TodoCreateOutput struct {
	Id int
}

// TodoUpdateInput 编辑
type TodoUpdateInput struct {
	Id int
	TodoOutputBase
}
type TodoUpdateOutput struct {
	Id int
}

// TodoDeleteInput 删除
type TodoDeleteInput struct {
	Id int
}
type TodoDeleteOutput struct{}
