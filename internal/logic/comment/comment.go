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

// GetCommentCount 评论计数
func (s *sComment) GetCommentCount(ctx context.Context) (count int64) {
	count, err := dao.Comment.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}
