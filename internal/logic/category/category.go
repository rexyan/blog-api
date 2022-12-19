package category

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *sCategory) CreateCategory(ctx context.Context, CategoryName string) (err error) {
	_, err = dao.Category.Ctx(ctx).Data(dao.Category.Columns().CategoryName, CategoryName).Insert()
	return
}

func (s *sCategory) GetCategoryPageList(ctx context.Context, page model.PageInput) (res *model.GteCategoryListOutPut, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GteCategoryListOutPut{}
	query := dao.Category.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GteCategoryListOutPut{}, err
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

func (s *sCategory) UpdateCategory(ctx context.Context, CategoryId int, CategoryName string) (err error) {
	_, err = dao.Category.Ctx(ctx).Data(dao.Category.Columns().CategoryName, CategoryName).Where(dao.Category.Columns().Id, CategoryId).Update()
	return
}

func (s *sCategory) DeleteCategory(ctx context.Context, CategoryId int) (err error) {
	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, CategoryId).Delete()
	return
}
