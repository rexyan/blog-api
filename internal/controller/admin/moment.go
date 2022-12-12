package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

type cMoment struct{}

var Moment = cMoment{}

func (c *cMoment) CreateMoment(ctx context.Context, req *admin.CreateMomentReq) (res *admin.CreateMomentRes, err error) {
	err = service.Moment().CreateMoment(ctx, model.CreateMomentInput{
		Content:    req.Content,
		CreateTime: gtime.New(req.CreateTime),
		Likes:      req.Likes,
		Published:  req.Published,
	})
	return
}
