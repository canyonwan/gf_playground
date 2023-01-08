package backend

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var Role = cRole{}

type cRole struct{}

// RoleCreateReq 创建角色

func (cr *cRole) Create(ctx context.Context, req *backend.RoleCreateReq) (res *backend.RoleCreateRes, err error) {
	create, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleBase: model.RoleBase{
			RoleName: req.RoleName,
			RoleDesc: req.RoleDesc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleCreateRes{Id: create.Id}, nil
}

// Update 更新角色
func (cr *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	create, err := service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleBase: model.RoleBase{
			RoleName: req.RoleName,
			RoleDesc: req.RoleDesc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleUpdateRes{Id: create.Id}, nil
}

// Delete 删除角色
func (cr *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	output, err := service.Role().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.RoleDeleteRes{Id: output.Id}, nil
}

// Page 角色分页
func (cr *cRole) Page(ctx context.Context, req *backend.RolePageReq) (res *backend.RolePageRes, err error) {
	out, err := service.Role().Page(ctx, model.RolePageInput{
		PageCommonReq: common.PageCommonReq{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RolePageRes{
		PageCommonRes: common.PageCommonRes{
			Content: out.Content,
			Page:    out.Page,
			Size:    out.Size,
			Total:   out.Total,
		},
	}, nil
}
