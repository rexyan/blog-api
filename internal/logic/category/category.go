package category

import (
	"blog-api/internal/dao"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) GetCategoryList(ctx context.Context) (res []*entity.Category, err error) {
	err = dao.Category.Ctx(ctx).Scan(&res)
	return
}
