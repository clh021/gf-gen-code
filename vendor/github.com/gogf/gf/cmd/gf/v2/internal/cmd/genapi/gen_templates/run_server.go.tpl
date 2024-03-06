package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
)

func runServe(ctx context.Context) {
	s := g.Server()

	// session
	s.SetSessionMaxAge(time.Hour * 24 * 30)
	s.SetSessionStorage(gsession.NewStorageMemory())

	// 静态资源
	s.AddStaticPath("/", "resource/public/dist") // 支持直接设置 gre 中的路径
	s.AddStaticPath("/assets", "resource/public/dist/assets") // 支持直接设置 gre 中的路径

	// 开发模式下，开启 openapi
	if IsDevelop() {
		runOpenApi(s)
	}

	// 解决 vue 静态资源 hashRoute 刷新问题
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.RedirectTo("/")
	})

	// 项目接口
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			hello.NewV1(),
		)
	})
	s.Run()
}