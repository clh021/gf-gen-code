package cmd

import "github.com/gogf/gf/v2/net/ghttp"

func runOpenApi(s *ghttp.Server) {
	s.SetOpenApiPath("/api.json")
	s.AddSearchPath("resource/public/resource") // 支持直接设置 gre 中的路径
	s.BindHandler("/apidoc", func(r *ghttp.Request) {
		r.Response.WriteTpl("/apidoc.html")
	})
}