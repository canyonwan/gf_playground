package account

import (
	"context"
	"gf_playground/api/v1/common"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}

func (sa *sAccount) Page(ctx context.Context, in model.AccountPageInput) (out *model.AccountPageOutput, err error) {
	var (
		m = dao.AccountInfo.Ctx(ctx)
	)
	out = &model.AccountPageOutput{
		PageCommonRes: common.PageCommonRes{
			Page: in.Page,
			Size: in.Size,
		},
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	var list []*entity.AccountInfo
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

func (sa *sAccount) Create(ctx context.Context, in model.AccountCreateInput) (out *model.AccountCreateOutput, err error) {
	id, err := dao.AccountInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.AccountCreateOutput{Id: int(id)}, nil
}

func (sa *sAccount) Update(ctx context.Context, in model.AccountUpdateInput) (out *model.AccountUpdateOutput, err error) {
	_, err = dao.AccountInfo.Ctx(ctx).Data(in).FieldsEx(dao.AccountInfo.Columns().Id).Where(g.Map{
		dao.AccountInfo.Columns().Id: in.Id,
	}).Update()
	if err != nil {
		return nil, err
	}
	return &model.AccountUpdateOutput{Id: in.Id}, nil
}

func (sa *sAccount) Delete(ctx context.Context, in model.AccountDeleteInput) (out *model.AccountDeleteOutput, err error) {
	_, err = dao.AccountInfo.Ctx(ctx).Where(g.Map{
		dao.AccountInfo.Columns().Id: in.Id,
	}).Unscoped().Delete()
	if err != nil {
		return nil, err
	}
	return &model.AccountDeleteOutput{Id: in.Id}, nil
}
