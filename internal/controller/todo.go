package controller

import (
	"context"
	"gf_playground/api/v1/admin"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	CTodo = cTodo{}
)

type cTodo struct{}

func (ct *cTodo) Get(ctx context.Context, req *admin.TodoGetReq) (res *admin.TodoGetRes, err error) {
	out, err := service.Todo().Get(ctx)
	if err != nil {
		return nil, err
	}
	return &admin.TodoGetRes{
		Id: out.Id,
		TodoResBase: admin.TodoResBase{
			Title:   out.Title,
			Content: out.Content,
		},
	}, err
}

func (ct *cTodo) Create(ctx context.Context, req *admin.TodoCreateReq) (res *admin.TodoCreateRes, err error) {
	out, err := service.Todo().Create(ctx, model.TodoCreateInput{
		TodoItemBase: model.TodoItemBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.TodoCreateRes{Id: out.Id}, err
}

func (ct *cTodo) Update(ctx context.Context, req *admin.TodoUpdateReq) (res *admin.TodoUpdateRes, err error) {
	output, err := service.Todo().Update(ctx, model.TodoUpdateInput{
		Id: req.Id,
		TodoItemBase: model.TodoItemBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.TodoUpdateRes{Id: output.Id}, err
}

func (ct *cTodo) Delete(ctx context.Context, req *admin.TodoDeleteReq) (res *admin.TodoDeleteRes, err error) {
	err = service.Todo().Delete(ctx, model.TodoDeleteInput{Id: req.Id})
	return
}
