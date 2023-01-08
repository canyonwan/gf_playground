package middleware

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/consts"
	"gf_playground/internal/dao"
	"gf_playground/internal/model/entity"
	"gf_playground/utility"
	"gf_playground/utility/response"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

func LoginBeforeAuth(r *ghttp.Request) (string, interface{}) {
	account := r.Get("account").String()
	pwd := r.Get("password").String()
	ctx := context.TODO()
	if account == "" || pwd == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}

	// 验证你输入的账号密码是否正确
	accountInfo := entity.AccountInfo{}
	err := dao.AccountInfo.Ctx(ctx).Where("account", account).Scan(&accountInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	// 如果通过帐号查到了, 说明加密盐也能查到
	// if 加密后的不等于DB里的密码,则不对
	if utility.EncryptPassword(pwd, accountInfo.UserSalt) != accountInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}

	return consts.GTokenAdminPrefix + strconv.Itoa(accountInfo.Id), accountInfo
}

// LoginAfterAuth 登录之后的操作
func LoginAfterAuth(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
	} else {
		respData.Code = 1
		// 获得admin登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		// 根据adminId获取用户信息
		var adminInfo entity.AccountInfo
		if err := dao.AccountInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo); err != nil {
			return
		}
		//	先通过角色查询权限 TODO: 完成角色和权限接口
		data := &backend.LoginRes{
			Type:         "Bearer",
			Token:        respData.GetString("token"),
			ExpireIn:     10 * 24 * 60 * 60, //单位秒,todo 根据实际情况修改
			RoleIds:      adminInfo.RoleIds,
			IsSuperAdmin: adminInfo.IsSuperAdmin,
		}
		response.JsonExit(r, respData.Code, respData.Msg, data)
	}
}

// AuthAfterFunc Auth通过后
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.AccountInfo
	if err := gconv.Struct(respData.GetString("data"), &adminInfo); err != nil {
		response.Auth(r)
		return
	}
	// 没有该帐号或已被拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminAccount, adminInfo.Account)
	r.SetCtxVar(consts.CtxAdminIsSuperAdmin, adminInfo.IsSuperAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
