package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cAbout struct{}

var About = cAbout{}

func (c *cAbout) About(ctx context.Context, req *admin.GetAboutReq) (res *admin.GetAboutRes, err error) {
	about, err := service.About().GetAbout(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(about, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cAbout) UpdateAbout(ctx context.Context, req *admin.UpdateAboutReq) (res *admin.UpdateAboutRes, err error) {
	err = service.About().UpdateAbout(ctx, &model.UpdateAboutInput{
		AboutItem: model.AboutItem{
			CommentEnabled: req.CommentEnabled,
			Content:        req.Content,
			MusicId:        req.MusicID,
			Title:          req.Title,
		},
	})
	return
}
