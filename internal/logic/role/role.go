package role

import (
	"context"
	"gf_playground/api/v1/common"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(&sRole{})
}

func (sr *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	id, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.RoleCreateOutput{}, err
	}
	return model.RoleCreateOutput{Id: int(id)}, nil
}

func (sr *sRole) Update(ctx context.Context, in model.RoleUpdateInput) (out model.RoleUpdateOutput, err error) {
	_, err = dao.RoleInfo.Ctx(ctx).Data(in).FieldsEx(dao.RoleInfo.Columns().Id).WherePri(in.Id).Update()
	if err != nil {
		return model.RoleUpdateOutput{}, err
	}
	return model.RoleUpdateOutput{Id: in.Id}, nil

}

func (sr *sRole) Delete(ctx context.Context, id int) (out model.RoleDeleteOutput, err error) {
	_, err = dao.RoleInfo.Ctx(ctx).WherePri(id).Unscoped().Delete()
	if err != nil {
		return model.RoleDeleteOutput{}, err
	}
	return model.RoleDeleteOutput{Id: id}, nil
}

func (sr *sRole) Page(ctx context.Context, in model.RolePageInput) (out *model.RolePageOutput, err error) {
	var m = dao.RoleInfo.Ctx(ctx)
	out = &model.RolePageOutput{
		PageCommonRes: common.PageCommonRes{
			Page: in.Page,
			Size: in.Size,
		},
	}
	// 查询分页数据
	listModel := m.Page(in.Page, in.Size)
	var list []*entity.RoleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}

	if len(list) == 0 {
		return out, nil
	}

	out.Total, err = m.Count()
	if err := listModel.Scan(&out.Content); err != nil {
		return out, err
	}

	return
}
