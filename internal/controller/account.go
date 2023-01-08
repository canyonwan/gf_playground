package controller

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	Account = cAccount{}
)

type cAccount struct{}

func (ca *cAccount) Page(ctx context.Context, req *backend.AccountPageReq) (res *backend.AccountPageRes, err error) {
	out, err := service.Account().Page(ctx, model.AccountPageInput{
		PageCommonReq: common.PageCommonReq{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AccountPageRes{
		PageCommonRes: common.PageCommonRes{
			Content: out.Content,
			Page:    out.Page,
			Size:    out.Size,
			Total:   out.Total,
		},
	}, nil
}

func (ca *cAccount) Create(ctx context.Context, req *backend.AccountCreateReq) (res *backend.AccountCreateRes, err error) {
	out, err := service.Account().Create(ctx, model.AccountCreateInput{
		AccountBase: model.AccountBase{
			Account:      req.Account,
			Password:     req.Password,
			RoleIds:      req.RoleIds,
			IsSuperAdmin: req.IsSuperAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AccountCreateRes{Id: out.Id}, nil
}

func (ca *cAccount) Update(ctx context.Context, req *backend.AccountUpdateReq) (res *backend.AccountUpdateRes, err error) {
	out, err := service.Account().Update(ctx, model.AccountUpdateInput{
		Id: req.Id,
		AccountBase: model.AccountBase{
			Account:      req.Account,
			Password:     req.Password,
			RoleIds:      req.RoleIds,
			IsSuperAdmin: req.IsSuperAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AccountUpdateRes{Id: out.Id}, nil
}

func (ca *cAccount) Delete(ctx context.Context, req *backend.AccountDeleteReq) (res *backend.AccountDeleteRes, err error) {
	out, err := service.Account().Delete(ctx, model.AccountDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AccountDeleteRes{Id: out.Id}, nil
}

func (ca *cAccount) Info(ctx context.Context, req *backend.AccountInfoReq) (res *backend.AccountInfoRes, err error) {
	info, err := service.Account().Info(ctx, model.AccountInfoInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AccountInfoRes{
		Id:           info.Id,
		Account:      info.Account,
		IsSuperAdmin: info.IsSuperAdmin,
		RoleIds:      info.RoleIds,
	}, nil
}
