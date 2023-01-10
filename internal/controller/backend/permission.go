package backend

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var Permission = cPermission{}

type cPermission struct{}

func (*cPermission) Create(ctx context.Context, req *backend.PermissionCreateReq) (res *backend.PermissionCreateRes, err error) {
	out, err := service.Permission().Create(ctx, model.PermissionCreateInput{
		PermissionBaseInput: model.PermissionBaseInput{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PermissionCreateRes{Id: out.Id}, nil
}

func (*cPermission) Update(ctx context.Context, req *backend.PermissionUpdateReq) (res *backend.PermissionUpdateRes, err error) {
	out, err := service.Permission().Update(ctx, model.PermissionUpdateInput{
		Id: req.Id,
		PermissionBaseInput: model.PermissionBaseInput{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PermissionUpdateRes{Id: out.Id}, nil
}

func (*cPermission) Delete(ctx context.Context, req *backend.PermissionDeleteReq) (res *backend.PermissionDeleteRes, err error) {
	out, err := service.Permission().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.PermissionDeleteRes{Id: out.Id}, nil
}

func (*cPermission) Page(ctx context.Context, req *backend.PermissionPageReq) (res *backend.PermissionPageRes, err error) {
	return
}

func (*cPermission) List(ctx context.Context, req *backend.PermissionListReq) (res *backend.PermissionListRes, err error) {
	return
}
