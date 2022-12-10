package cmd

import (
	"blog-api/internal/controller/admin"
	"blog-api/internal/service"
	"context"

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
			s := g.Server()
			//全局中间件注册
			s.Use(service.Middleware().CustomResponseHandler, ghttp.MiddlewareCORS)
			// 后台路由
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				// 不需要 JWT Token 校验的路由
				group.GET("/login", admin.User.Login)

				// 需要  JWT Token 的路由
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Jwt)
					group.Bind(admin.Dashboard.Index)
					group.Bind(admin.Blog.CreateBlog)
					group.Bind(admin.Blog.UpdateBlog)
					group.Bind(admin.Blog.GetBlogList)
					group.Bind(admin.Blog.GetBlogDetail)
					group.Bind(admin.Category.CategoryAndTag)
				})
			})
			s.Run()
			return nil
		},
	}
)
