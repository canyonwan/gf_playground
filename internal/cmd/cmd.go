package cmd

import (
	"context"
	"gf_playground/internal/consts"
	"gf_playground/internal/controller/backend"
	"gf_playground/internal/controller/frontend"
	"gf_playground/internal/middleware"
	"gf_playground/internal/service"
	"github.com/goflyfox/gtoken/gtoken"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
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
			//oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// 开启gfToken
			gfAdminToken := gtoken.GfToken{
				ServerName:       "gf_admin_auth",
				CacheMode:        2,
				MultiLogin:       false,
				LoginPath:        consts.AdminLoginPath,
				LoginBeforeFunc:  middleware.Login,
				LoginAfterFunc:   middleware.LoginAfterAuth,
				LogoutPath:       consts.AdminLogoutPath,
				AuthFailMsg:      consts.AuthFailMsg,
				AuthAfterFunc:    middleware.AuthAfterFunc,
				AuthExcludePaths: g.SliceStr{"/v1/backend/login"}, // 不拦截白名单
			}

			// 小程序前端 gfToken
			gfFrontendToken := gtoken.GfToken{
				ServerName:       "gf_frontend_auth",
				CacheMode:        2,
				MultiLogin:       false,
				LoginPath:        consts.LoginPath,
				LoginBeforeFunc:  middleware.Login,
				LoginAfterFunc:   middleware.LoginAfterAuth,
				LogoutPath:       consts.LogoutPath,
				AuthFailMsg:      consts.AuthFailMsg,
				AuthAfterFunc:    middleware.AuthAfterFunc,
				AuthExcludePaths: g.SliceStr{"/v1/frontend/login"}, // 不拦截白名单
			}

			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareCORS,
					//service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)

				// 后台管理 全部都需要登录的
				group.Group("/backend", func(backendGroup *ghttp.RouterGroup) {
					// 使用gfToken中间件
					if err := gfAdminToken.Middleware(ctx, backendGroup); err != nil {
						panic(err)
					}

					backendGroup.Bind(
						backend.Banner,
						backend.DataStatistics,
						backend.Position,
						backend.Account,
						backend.Role,
						backend.Permission,
						backend.File,
					)
				})

				// 小程序前端
				group.Group("/frontend", func(frontendGroup *ghttp.RouterGroup) {
					frontendGroup.Bind(frontend.BannerFrontend)
					// 使用gfToken中间件
					if err := gfFrontendToken.Middleware(ctx, frontendGroup); err != nil {
						panic(err)
					}
					frontendGroup.Bind(
						frontend.Todo,
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
