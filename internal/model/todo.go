package model

import "github.com/gogf/gf/v2/os/gtime"

type TodoItemBase struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
	DeletedAt gtime.Time `json:"deleted_at"`
}
type TodoOutputBase struct {
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
	DeletedAt gtime.Time `json:"deleted_at"`
}

// TodoPageGetInput 获取
type TodoPageGetInput struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
type TodoPageGetOutput struct {
	Page    int            `json:"page"`
	Size    int            `json:"size"`
	Content []TodoItemBase `json:"content"`
	Total   int            `json:"total"`
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
	TodoItemBase
}
type TodoUpdateOutput struct{}

// TodoDeleteInput 删除
type TodoDeleteInput struct {
	Id int
}
type TodoDeleteOutput struct{}
