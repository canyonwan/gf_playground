package login

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"gf_playground/utility"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(&sLogin{})
}

// Login backend 用户登录就是AccountInfo登录
func (sl *sLogin) Login(ctx context.Context, in model.LoginInput) error {
	// 验证你输入的账号密码是否正确
	accountInfo := entity.AccountInfo{}
	err := dao.AccountInfo.Ctx(ctx).Where("account", in.Account).Scan(&accountInfo)
	if err != nil {
		return err
	}
	// 如果通过帐号查到了, 说明加密盐也能查到
	// if 加密后的不等于DB里的密码,则不对
	if utility.EncryptPassword(in.Password, accountInfo.UserSalt) != accountInfo.Password {
		return gerror.New("帐号或密码错误")
	}
	if err := service.Session().SetUser(ctx, &accountInfo); err != nil {
		return err
	}

	return nil
}
