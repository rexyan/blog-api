package visitor

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"time"
)

type sVisitor struct{}

func init() {
	service.RegisterVisitor(New())
}
func New() *sVisitor {
	return &sVisitor{}
}

func (s *sVisitor) GetTodayPvUv(ctx context.Context) (res *model.TodayPvUvOutput, err error) {
	dateStr := time.Now().Format("01-02")
	visitRecordCls := dao.VisitRecord.Columns()
	err = dao.VisitRecord.Ctx(ctx).
		Where(visitRecordCls.Date, dateStr).
		Scan(&res)
	return
}

func (s *sVisitor) GetLatestMonthPvUv(ctx context.Context) (res []*model.LatestMonthPvUvOutput, err error) {
	err = dao.VisitRecord.Ctx(ctx).
		Limit(30).
		Scan(&res)
	return
}

func (s *sVisitor) GetCityVisitor(ctx context.Context) (res []*model.CityVisitorOutput, err error) {
	err = dao.CityVisitor.Ctx(ctx).Scan(&res)
	return
}
