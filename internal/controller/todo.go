package controller

import (
	"context"
	"gf_playground/api/v1/app"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	CTodo = cTodo{}
)

type cTodo struct{}

func (ct *cTodo) GetList(ctx context.Context, req *app.TodoListGetReq) (res *app.TodoListGetRes, err error) {
	//res = &app.TodoListGetRes{}
	out, err := service.Todo().GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &app.TodoListGetRes{List: out.List}, nil
}

func (ct *cTodo) GetPage(ctx context.Context, req *app.TodoPageGetReq) (res *app.TodoPageGetRes, err error) {
	out, err := service.Todo().GetPage(ctx, model.TodoPageGetInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &app.TodoPageGetRes{
		PageCommonRes: common.PageCommonRes{
			Page:    out.Page,
			Size:    out.Size,
			Content: out.Content,
			Total:   out.Total,
		},
	}, nil
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
		TodoItemBase: model.TodoItemBase{
			Id:      req.Id,
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
