package controller

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var Login = cLogin{}

type cLogin struct{}

func (l *cLogin) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	res = &backend.LoginRes{}
	err = service.Login().Login(ctx, model.LoginInput{Account: req.Account, Password: req.Password})
	if err != nil {
		return nil, err
	}

	var accountInfo *entity.AccountInfo
	accountInfo = service.Session().GetUser(ctx)
	res.Info = g.Map{
		"id":           accountInfo.Id,
		"account":      accountInfo.Account,
		"isSuperAdmin": accountInfo.IsSuperAdmin,
		"roleIds":      accountInfo.RoleIds,
	}
	return
}
