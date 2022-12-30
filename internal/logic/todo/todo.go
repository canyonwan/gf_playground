package todo

import (
	"context"
	"fmt"
	"gf_playground/api/v1/common"
	"gf_playground/internal/dao"
	model "gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type sTodo struct{}

func init() {
	service.RegisterTodo(&sTodo{})
}

func (s *sTodo) GetList(ctx context.Context) (out *model.TodoListGetOutput, err error) {
	// 先实例化地址 否则会报空指针
	// todo 为什么会报空指针
	out = &model.TodoListGetOutput{}
	err = dao.Todo.Ctx(ctx).OrderAsc(dao.Todo.Columns().Id).Scan(&out.List)
	return

	//return
	// 先实例化地址 否则会报空指针
	// todo 为什么会报空指针
	//out = &model.TodoListGetOutput{}
	//err = dao.Todo.Ctx(ctx).OrderAsc(dao.Todo.Columns().Id).Scan(&out.List)
	//if err != nil {
	//	return nil, err
	//}
	//return out, nil
}

func (s *sTodo) GetPage(ctx context.Context, in model.TodoPageGetInput) (out *model.TodoPageGetOutput, err error) {
	// m = *gdb.Model 获取数据库操作对象
	var m = dao.Todo.Ctx(ctx)
	out = &model.TodoPageGetOutput{
		PageCommonRes: common.PageCommonRes{
			Page: in.Page,
			Size: in.Size,
		},
	}
	listModel := m.Page(in.Page, in.Size)

	var list []*entity.Todo // 定义TodoItem模型
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	//
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	// 有数据 获取总数
	out.Total, err = m.Count()
	fmt.Println("out", out)
	if err := listModel.Scan(&out.Content); err != nil {
		return out, err
	}
	return
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
