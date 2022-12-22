package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var Visitor = cVisitor{}

type cVisitor struct{}

func (c *cVisitor) GetVisitorList(ctx context.Context, req *admin.GetVisitorListReq) (res *admin.GetVisitorListRes, err error) {
	visitorList, err := service.Visitor().GetVisitorList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(visitorList, &res)
	return
}

func (c *cVisitor) DeleteVisitor(ctx context.Context, req *admin.DeleteVisitorReq) (res *admin.DeleteVisitorRes, err error) {
	err = service.Visitor().DeleteVisitor(ctx, req.VisitorId)
	return
}
