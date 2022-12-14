package account

import (
	"context"
	"gf_playground/api/v1/common"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"gf_playground/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
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

func (sa *sAccount) Create(ctx context.Context, in model.AccountCreateInput) (out model.AccountCreateOutput, err error) {
	userSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	in.UserSalt = userSalt
	id, err := dao.AccountInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AccountCreateOutput{Id: int(id)}, nil
}

func (sa *sAccount) Update(ctx context.Context, in model.AccountUpdateInput) (out *model.AccountUpdateOutput, err error) {
	// 注: 如果用户修改了密码, 需要重新加密;没有则反之
	// 1. 判断是否修改了密码
	if in.Password != "" {
		userSalt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, userSalt)
		in.UserSalt = userSalt
	}
	_, err = dao.AccountInfo.Ctx(ctx).Data(in).FieldsEx(dao.AccountInfo.Columns().Id).Where(g.Map{
		dao.AccountInfo.Columns().Id: in.Id,
	}).OmitEmpty().Update()
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

func (sa *sAccount) Info(ctx context.Context, in model.AccountInfoInput) (out *model.AccountInfoOutput, err error) {
	out = &model.AccountInfoOutput{}
	if err = dao.AccountInfo.Ctx(ctx).Where(dao.AccountInfo.Columns().Id, in.Id).Scan(out); err != nil {
		return nil, err
	}
	return
}
