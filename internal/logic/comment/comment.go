package comment

import (
	"blog-api/internal/dao"
	"blog-api/internal/service"
	"context"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) GetCommentCount(ctx context.Context) (count int64, err error) {
	count, err = dao.Comment.Ctx(ctx).Count()
	return
}
