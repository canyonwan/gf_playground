package todo

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

type sTodo struct{}

func init() {
	service.RegisterTodo(&sTodo{})
}

func (s *sTodo) Get(ctx context.Context) (out model.TodoGetOutput, err error) {
	return model.TodoGetOutput{
		Id: out.Id,
		TodoOutputBase: model.TodoOutputBase{
			Title:   out.Title,
			Content: out.Content,
		},
	}, err

}

func (s *sTodo) Create(ctx context.Context, in model.TodoCreateInput) (out model.TodoCreateOutput, err error) {
	id, err := dao.Todo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.TodoCreateOutput{}, err
	}

	return model.TodoCreateOutput{Id: int(id)}, err
}

func (s *sTodo) Update(ctx context.Context, in model.TodoUpdateInput) (out model.TodoUpdateOutput, err error) {
	id, err := dao.Todo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.TodoUpdateOutput{}, err
	}
	return model.TodoUpdateOutput{Id: int(id)}, err
}

func (s *sTodo) Delete(ctx context.Context, in model.TodoDeleteInput) (err error) {
	//dao.Todo.Ctx(ctx).Where(g.Map{
	//	dao.Todo.Columns().Id: in.Id,
	//}).Delete()
	_, err = dao.Todo.Ctx(ctx).WherePri(in.Id).Delete()
	if err != nil {
		return err
	}
	return err
}
