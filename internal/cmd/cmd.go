package cmd

import (
	"context"

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

			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(ghttp.MiddlewareCORS)
				group.Bind(
					controller.Banner,
					controller.DataStatistics,
					controller.CTodo,
				)
			})
			s.Run()
			return nil
		},
	}
)
