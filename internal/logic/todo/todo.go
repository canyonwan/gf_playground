package todo

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type sTodo struct{}

func init() {
	service.RegisterTodo(&sTodo{})
}

func (s *sTodo) Get(context.Context) (out model.TodoGetOutput, err error) {
	return model.TodoGetOutput{
		Id: out.Id,
		TodoOutputBase: model.TodoOutputBase{
			Title:   out.Title,
			Content: out.Content,
		},
	}, err

}

func (s *sTodo) Create(ctx context.Context, in *model.TodoCreateInput) (out *model.TodoCreateOutput, err error) {
	id, err := dao.Todo.Ctx(ctx).Data(&in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.TodoCreateOutput{Id: int(id)}, err
}

func (s *sTodo) Update(ctx context.Context, in model.TodoUpdateInput) (err error) {
	_, err = dao.Todo.Ctx(ctx).Data(in).FieldsEx(dao.Todo.Columns().Id).Where(dao.Todo.Columns().Id, in.Id).Update()
	if err != nil {
		return err
	}
	return err
}

func (s *sTodo) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.Todo.Ctx(ctx).Where(g.Map{
		dao.Todo.Columns().Id: id,
	}).Unscoped().Delete()
	if err != nil {
		return err
	}
	return nil
	//_, err = dao.Todo.Ctx(ctx).WherePri(in.Id).Delete()
	//if err != nil {
	//	return err
	//}
	//return err
}
