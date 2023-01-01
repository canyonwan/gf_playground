package controller

import (
	"context"
	"gf_playground/api/v1/admin"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
)

var (
	Account = cAccount{}
)

type cAccount struct{}

func (ca *cAccount) Page(ctx context.Context, req *admin.AccountPageReq) (res *admin.AccountPageRes, err error) {
	out, err := service.Account().Page(ctx, model.AccountPageInput{
		PageCommonReq: common.PageCommonReq{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.AccountPageRes{
		PageCommonRes: common.PageCommonRes{
			Content: out.Content,
			Page:    out.Page,
			Size:    out.Size,
			Total:   out.Total,
		},
	}, nil
}

func (ca *cAccount) Create(ctx context.Context, req *admin.AccountCreateReq) (res *admin.AccountCreateRes, err error) {
	out, err := service.Account().Create(ctx, model.AccountCreateInput{
		AccountBase: model.AccountBase{
			Account:      req.Account,
			Password:     req.Password,
			IsSuperAdmin: req.IsSuperAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.AccountCreateRes{Id: out.Id}, nil
}

func (ca *cAccount) Update(ctx context.Context, req *admin.AccountUpdateReq) (res *admin.AccountUpdateRes, err error) {
	out, err := service.Account().Update(ctx, model.AccountUpdateInput{
		AccountInfo: entity.AccountInfo{
			Id:           req.Id,
			Account:      req.Account,
			Password:     req.Password,
			IsSuperAdmin: req.IsSuperAdmin,
			RoleIds:      req.RoleIds,
			UserSalt:     req.UserSalt,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.AccountUpdateRes{Id: out.Id}, nil
}

func (ca *cAccount) Delete(ctx context.Context, req *admin.AccountDeleteReq) (res *admin.AccountDeleteRes, err error) {
	out, err := service.Account().Delete(ctx, model.AccountDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &admin.AccountDeleteRes{Id: out.Id}, nil
}
