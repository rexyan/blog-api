package tag

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTag struct{}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

// GetTagList Tag 列表
func (s *sTag) GetTagList(ctx context.Context) (tags []*entity.Tag, err error) {
	err = dao.Tag.Ctx(ctx).Scan(&tags)
	return
}

// GetTagIdsByBlogId 获取文章关联的 Tag ID 列表
func (s *sTag) GetTagIdsByBlogId(ctx context.Context, BlogId int) (tags []*entity.BlogTag, err error) {
	err = dao.BlogTag.Ctx(ctx).Where(dao.BlogTag.Columns().BlogId, BlogId).Scan(&tags)
	return
}

// DashboardTagStatistics 看板 Tag 分析数据
func (s *sTag) DashboardTagStatistics(ctx context.Context) (res *model.DashboardTagOutput, err error) {
	res = &model.DashboardTagOutput{}
	tagList, err := s.GetTagList(ctx)
	if err != nil {
		return nil, err
	}
	for _, data := range tagList {
		res.Legend = append(res.Legend, data.TagName)
		// 获取标签关联的文章的计数
		blogCount := service.Blog().GetBlogCountByTagId(ctx, data.Id)
		res.Series = append(res.Series, model.Series{
			ID:    gconv.Int(data.Id),
			Name:  data.TagName,
			Value: gconv.Int(blogCount),
		})
	}
	return
}

func (s *sTag) CreateTag(ctx context.Context, in *model.CreateTagInput) (err error) {
	_, err = dao.Tag.Ctx(ctx).Data(g.Map{
		dao.Tag.Columns().Color:   in.Color,
		dao.Tag.Columns().TagName: in.TagName,
	}).Insert()
	return
}

func (s *sTag) GetTagPageList(ctx context.Context, page model.PageInput) (res *model.GetTagListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetTagListOutput{}
	query := dao.Tag.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetTagListOutput{}, err
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

func (s *sTag) UpdateTag(ctx context.Context, in *model.UpdateTagInput) (err error) {
	_, err = dao.Tag.Ctx(ctx).Data(g.Map{
		dao.Tag.Columns().TagName: in.TagName,
		dao.Tag.Columns().Color:   in.Color,
	}).Where(dao.Tag.Columns().Id, in.ID).Update()
	return
}

func (s *sTag) DeleteTag(ctx context.Context, TagId int) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().Id, TagId).Delete()
	return
}
