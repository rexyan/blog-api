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

// GetTodayPvUv 获取今天的 PvUv
func (s *sVisitor) GetTodayPvUv(ctx context.Context) (res *model.TodayPvUvOutput, err error) {
	dateStr := time.Now().Format("01-02")
	visitRecordCls := dao.VisitRecord.Columns()
	err = dao.VisitRecord.Ctx(ctx).
		Where(visitRecordCls.Date, dateStr).
		Scan(&res)
	return
}

// DashBoardMonthPvUvStatistics 看板月度 PvUv 分析
func (s *sVisitor) DashBoardMonthPvUvStatistics(ctx context.Context) (res *model.DashBoardMonthPvUv, err error) {
	res = &model.DashBoardMonthPvUv{}
	latestMonthPvUv, err := s.GetLatestMonthPvUv(ctx)
	if err != nil {
		return nil, err
	}
	for _, data := range latestMonthPvUv {
		res.Uv = append(res.Uv, data.Uv)
		res.Pv = append(res.Pv, data.Pv)
		res.Date = append(res.Date, data.Date)
	}
	return
}

// GetLatestMonthPvUv 获取最近一个月的 PvUv 数据
func (s *sVisitor) GetLatestMonthPvUv(ctx context.Context) (res []*model.LatestMonthPvUvOutput, err error) {
	err = dao.VisitRecord.Ctx(ctx).
		Limit(30).
		Scan(&res)
	return
}

// GetCityVisitor 获取城市访问数据
func (s *sVisitor) GetCityVisitor(ctx context.Context) (res []*model.CityVisitorOutput, err error) {
	err = dao.CityVisitor.Ctx(ctx).Scan(&res)
	return
}
