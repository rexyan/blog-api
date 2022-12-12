package moment

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
		momentCls.IsPublished: input.Published,
	}).Insert()
	return
}
