package cmd

import (
	"context"
	"gf_playground/internal/consts"
	"gf_playground/internal/service"
	"github.com/goflyfox/gtoken/gtoken"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_playground/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)
			// OpenApi自定义信
			oai.Info.Title = `Canyonwan的接口文档`
			//oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`
			loginFunc := loginAuth

			// 开启gfToken
			gfToken := gtoken.GfToken{
				ServerName:       "gf_shop_demo",
				MultiLogin:       true,
				LoginPath:        consts.LoginPath,
				LogoutPath:       consts.LogoutPath,
				LoginBeforeFunc:  loginFunc,
				AuthExcludePaths: g.SliceStr{"/backend/login", "/backend/logout"}, // 不拦截白名单
			}

			s.Group("/v1", func(group *ghttp.RouterGroup) {
				// 使用gfToken中间件
				if err := gfToken.Middleware(ctx, group); err != nil {
					panic(err)
				}
				group.Middleware(
					ghttp.MiddlewareCORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Banner,
					controller.DataStatistics,
					controller.Todo,
					controller.Position,
					controller.Account,
					//controller.Login,
				)
			})
			s.Run()
			return nil
		},
	}
)

func loginAuth(r *ghttp.Request) (string, interface{}) {
	account := r.Get("account").String()
	pwd := r.Get("password").String()
	if account == "" || pwd == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	return account, "1"
}
