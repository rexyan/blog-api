package moment

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMoment struct{}

func init() {
	service.RegisterMoment(New())
}
func New() *sMoment {
	return &sMoment{}
}

func (s *sMoment) CreateMoment(ctx context.Context, input model.CreateMomentInput) (err error) {
	momentCls := dao.Moment.Columns()
	_, err = dao.Moment.Ctx(ctx).Data(g.Map{
		momentCls.Content:     input.Content,
		momentCls.CreateTime:  input.CreateTime,
		momentCls.Likes:       input.Likes,
		momentCls.IsPublished: input.IsPublished,
	}).Insert()
	return
}

func (s *sMoment) GetMomentList(ctx context.Context, page model.PageInput) (res *model.GetMomentListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetMomentListOutput{}
	query := dao.Moment.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetMomentListOutput{}, err
	}

	// 总数
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	// 页码信息
	pageInfo := r.GetPage(gconv.Int(count), page.PageSize)
	err = gconv.Scan(result, &res.List)
	if err != nil {
		return nil, err
	}

	service.Paginate().Paginate(&res.CommonPageHelper, count, page, result, pageInfo)
	return
}

func (s *sMoment) GetMomentDetail(ctx context.Context, MomentId int) (res *model.MomentDetailOutput, err error) {
	record, err := dao.Moment.Ctx(ctx).Where(dao.Moment.Columns().Id, MomentId).One()
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(record, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sMoment) UpdateMoment(ctx context.Context, in *model.UpdateMomentInput) error {
	momentCls := dao.Moment.Columns()
	_, err := dao.Moment.Ctx(ctx).Where(momentCls.Id, in.ID).Data(g.Map{
		momentCls.Content:     in.Content,
		momentCls.CreateTime:  in.CreateTime,
		momentCls.IsPublished: in.IsPublished,
		momentCls.Likes:       in.Likes,
	}).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *sMoment) UpdateMomentPublished(ctx context.Context, MomentId int, IsPublished bool) (err error) {
	_, err = dao.Moment.Ctx(ctx).Where(dao.Moment.Columns().Id, MomentId).Data(dao.Moment.Columns().IsPublished, IsPublished).Update()
	return
}

func (s *sMoment) DeleteMoment(ctx context.Context, MomentId int) (err error) {
	_, err = dao.Moment.Ctx(ctx).Where(dao.Moment.Columns().Id, MomentId).Delete()
	return
}
