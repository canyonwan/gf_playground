package controller

import (
	"context"
	"gf_playground/api/v1/common"
	"gf_playground/api/v1/frontend"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	Todo = cTodo{}
)

type cTodo struct{}

func (ct *cTodo) GetList(ctx context.Context, req *frontend.TodoListGetReq) (res *frontend.TodoListGetRes, err error) {
	//res = &frontend.TodoListGetRes{}
	out, err := service.Todo().GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &frontend.TodoListGetRes{List: out.List}, nil
}

func (ct *cTodo) GetPage(ctx context.Context, req *frontend.TodoPageGetReq) (res *frontend.TodoPageGetRes, err error) {
	out, err := service.Todo().GetPage(ctx, model.TodoPageGetInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.TodoPageGetRes{
		PageCommonRes: common.PageCommonRes{
			Page:    out.Page,
			Size:    out.Size,
			Content: out.Content,
			Total:   out.Total,
		},
	}, nil
}

func (ct *cTodo) Create(ctx context.Context, req *frontend.TodoCreateReq) (res *frontend.TodoCreateRes, err error) {
	out, err := service.Todo().Create(ctx, &model.TodoCreateInput{
		TodoItemOutputBase: model.TodoItemOutputBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.TodoCreateRes{Id: out.Id}, err
}

func (ct *cTodo) Update(ctx context.Context, req *frontend.TodoUpdateReq) (res *frontend.TodoUpdateRes, err error) {
	err = service.Todo().Update(ctx, model.TodoUpdateInput{
		Id: req.Id,
		TodoItemOutputBase: model.TodoItemOutputBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	return
}

func (ct *cTodo) Delete(ctx context.Context, req *frontend.TodoDeleteReq) (res *frontend.TodoDeleteRes, err error) {
	err = service.Todo().Delete(ctx, req.Id)
	return
}
