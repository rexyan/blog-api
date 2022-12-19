package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type cMoment struct{}

var Moment = cMoment{}

func (c *cMoment) CreateMoment(ctx context.Context, req *admin.CreateMomentReq) (res *admin.CreateMomentRes, err error) {
	err = service.Moment().CreateMoment(ctx, model.CreateMomentInput{
		Content:     req.Content,
		CreateTime:  gtime.New(req.CreateTime),
		Likes:       req.Likes,
		IsPublished: req.Published,
	})
	return
}

func (c *cMoment) GetMomentList(ctx context.Context, req *admin.GetMomentListReq) (res *admin.GetMomentListRes, err error) {
	res = &admin.GetMomentListRes{}
	momentList, err := service.Moment().GetMomentList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	err = gconv.Scan(momentList, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cMoment) GetMomentDetail(ctx context.Context, req *admin.MomentDetailReq) (res *admin.MomentDetailRes, err error) {
	moment, err := service.Moment().GetMomentDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(moment, &res)
	return
}

func (c *cMoment) UpdateMoment(ctx context.Context, req *admin.UpdateMomentReq) (res *admin.UpdateMomentRes, err error) {
	err = service.Moment().UpdateMoment(ctx, &model.UpdateMomentInput{
		ID: req.ID,
		CreateMomentInput: model.CreateMomentInput{
			Content:     req.Content,
			CreateTime:  gtime.New(req.CreateTime),
			Likes:       gconv.String(req.Likes),
			IsPublished: req.IsPublished,
		},
	})
	return
}

func (c *cMoment) UpdateMomentPublished(ctx context.Context, req *admin.UpdateMomentPublishedReq) (res *admin.UpdateMomentPublishedRes, err error) {
	err = service.Moment().UpdateMomentPublished(ctx, req.Id, req.Published)
	return
}
