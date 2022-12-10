package tag

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
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
