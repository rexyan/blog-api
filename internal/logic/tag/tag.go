package tag

import (
	"blog-api/internal/dao"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
)

type sTag struct{}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

func (s *sTag) GetTagList(ctx context.Context) (tags []*entity.Tag, err error) {
	err = dao.Tag.Ctx(ctx).Scan(&tags)
	return
}
