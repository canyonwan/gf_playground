package model

import (
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

// TodoItemBase Model层的json也要写 否则被controller层返回给客户端的时候将会是大写

type TodoItemOutputBase struct {
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

// TodoPageGetOutput 获取
type TodoPageGetOutput struct {
	common.PageCommonRes
}

// TodoListGetInput 获取
type TodoListGetInput struct{}

// TodoListGetOutput 获取
type TodoListGetOutput struct {
	List []*entity.Todo `json:"list"`
}

// TodoCreateInput 新增
type TodoCreateInput struct {
	TodoItemOutputBase
}
type TodoCreateOutput struct {
	Id int
}

// TodoUpdateInput 编辑
type TodoUpdateInput struct {
	Id int `json:"id"`
	TodoItemOutputBase
}
type TodoUpdateOutput struct{}

// TodoDeleteInput 删除
type TodoDeleteInput struct {
	Id int
}
type TodoDeleteOutput struct{}
