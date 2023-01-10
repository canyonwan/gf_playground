package permission

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(&sPermission{})
}

func (*sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out *model.PermissionCreateOutput, err error) {
	id, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.PermissionCreateOutput{Id: uint(id)}, nil
}

func (*sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) (out *model.PermissionUpdateOutput, err error) {
	_, err = dao.PermissionInfo.Ctx(ctx).Data(in).FieldsEx(in.Id).Update()
	if err != nil {
		return nil, err
	}
	return &model.PermissionUpdateOutput{Id: in.Id}, nil
}

func (*sPermission) Delete(ctx context.Context, id uint) (out *model.PermissionDeleteOutput, err error) {
	_, err = dao.PermissionInfo.Ctx(ctx).Where(dao.PermissionInfo.Columns().Id, id).Delete()
	if err != nil {
		return nil, err
	}
	return &model.PermissionDeleteOutput{Id: id}, nil
}

func (*sPermission) Page(ctx context.Context, in model.PermissionPageInput) (out *model.PermissionPageOutput, err error) {
	return
}

func (*sPermission) List(ctx context.Context, in model.PermissionListInput) (out *model.PermissionListOutput, err error) {
	return
}
