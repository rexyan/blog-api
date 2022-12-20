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
					group.Bind(admin.Blog.DeleteBlog)
					group.Bind(admin.Blog.GetBlogList)
					group.Bind(admin.Blog.GetBlogDetail)
					group.Bind(admin.Blog.UpdateBlogTop)
					group.Bind(admin.Blog.GetBlogIdAndTitle)
					group.Bind(admin.Blog.UpdateBlogRecommend)
					group.Bind(admin.Blog.UpdateBlogVisibility)
					group.Bind(admin.Moment.CreateMoment)
					group.Bind(admin.Moment.GetMomentList)
					group.Bind(admin.Moment.GetMomentDetail)
					group.Bind(admin.Moment.UpdateMoment)
					group.Bind(admin.Moment.UpdateMomentPublished)
					group.Bind(admin.Moment.DeleteMoment)
					group.Bind(admin.Category.CategoryAndTag)
					group.Bind(admin.Category.CreateCategory)
					group.Bind(admin.Category.GetCategoryList)
					group.Bind(admin.Category.UpdateCategory)
					group.Bind(admin.Category.DeleteCategory)
					group.Bind(admin.Tag.CreateTag)
					group.Bind(admin.Tag.GetTagList)
					group.Bind(admin.Tag.UpdateTag)
					group.Bind(admin.Tag.DeleteTag)
					group.Bind(admin.Comment.UpdateComment)
					group.Bind(admin.Comment.DeleteComment)
					group.Bind(admin.Comment.GetCommentTreeList)
					group.Bind(admin.Comment.UpdateCommentNotice)
					group.Bind(admin.Comment.UpdateCommentPublish)
					group.Bind(admin.SetSetting.GetSetSetting)
					group.Bind(admin.SetSetting.UpdateSetSetting)
				})
			})
			s.Run()
			return nil
		},
	}
)
