package controller

import (
	"context"
	"gf_playground/api/v1/app"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	CTodo = cTodo{}
)

type cTodo struct{}

func (ct *cTodo) Get(ctx context.Context, req *app.TodoGetReq) (res *app.TodoGetRes, err error) {
	out, err := service.Todo().Get(ctx)
	if err != nil {
		return nil, err
	}
	return &app.TodoGetRes{
		Id: out.Id,
		TodoResBase: app.TodoResBase{
			Title:   out.Title,
			Content: out.Content,
		},
	}, err
}

func (ct *cTodo) Create(ctx context.Context, req *app.TodoCreateReq) (res *app.TodoCreateRes, err error) {
	out, err := service.Todo().Create(ctx, &model.TodoCreateInput{
		TodoOutputBase: model.TodoOutputBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return &app.TodoCreateRes{Id: out.Id}, err
}

func (ct *cTodo) Update(ctx context.Context, req *app.TodoUpdateReq) (res *app.TodoUpdateRes, err error) {
	err = service.Todo().Update(ctx, model.TodoUpdateInput{
		Id: req.Id,
		TodoOutputBase: model.TodoOutputBase{
			Title:   req.Title,
			Content: req.Content,
		},
	})
	return
}

func (ct *cTodo) Delete(ctx context.Context, req *app.TodoDeleteReq) (res *app.TodoDeleteRes, err error) {
	err = service.Todo().Delete(ctx, req.Id)
	return
}
