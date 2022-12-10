package category

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// GetCategoryList 获取 Category 列表
func (s *sCategory) GetCategoryList(ctx context.Context) (res []*entity.Category, err error) {
	err = dao.Category.Ctx(ctx).Scan(&res)
	return
}

// DashboardCategoryStatistics 看板 Category 分析数据
func (s *sCategory) DashboardCategoryStatistics(ctx context.Context) (res *model.DashboardCategoryOutput, err error) {
	res = &model.DashboardCategoryOutput{}
	categoryList, err := s.GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	for _, data := range categoryList {
		res.Legend = append(res.Legend, data.CategoryName)
		// 获取 Category 关联的文章的计数
		blogCount := service.Blog().GetBlogCountByCategoryId(ctx, data.Id)
		res.Series = append(res.Series, model.Series{
			ID:    gconv.Int(data.Id),
			Name:  data.CategoryName,
			Value: gconv.Int(blogCount),
		})
	}
	return
}
