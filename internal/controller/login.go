package controller

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var Login = cLogin{}

type cLogin struct{}

func (l *cLogin) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	res = &backend.LoginRes{}
	err = service.Login().Login(ctx, model.LoginInput{Account: req.Account, Password: req.Password})
	if err != nil {
		return nil, err
	}

	res.Info = service.Session().GetUser(ctx)
	return
}
