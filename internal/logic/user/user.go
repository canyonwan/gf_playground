package user

import (
	"context"
	"gf_playground/internal/consts"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

func (*sUser) UserInfoGet(ctx context.Context) (out *model.UserInfoGetOutput, err error) {
	return &model.UserInfoGetOutput{
		Id:           gconv.Uint(ctx.Value(consts.CtxAdminId)),
		Account:      gconv.String(ctx.Value(consts.CtxAdminAccount)),
		Username:     gconv.String(ctx.Value(consts.CtxAdminUsername)),
		Gender:       gconv.Int8(ctx.Value(consts.CtxAdminGender)),
		Avatar:       gconv.String(ctx.Value(consts.CtxAdminAvatar)),
		Email:        gconv.String(ctx.Value(consts.CtxAdminEmail)),
		Phone:        gconv.String(ctx.Value(consts.CtxAdminPhone)),
		RoleIds:      gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
		Permissions:  gconv.SliceStr(ctx.Value(consts.CtxAdminPermissions)),
		IsSuperAdmin: gconv.Int8(ctx.Value(consts.CtxAdminIsSuperAdmin)),
	}, nil
}
