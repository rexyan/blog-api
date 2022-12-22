package visitor

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

func (s *sVisitor) GetVisitorList(ctx context.Context, page model.PageInput) (res *model.GetVisitListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetVisitListOutput{}
	query := dao.Visitor.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetVisitListOutput{}, err
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

func (s *sVisitor) DeleteVisitor(ctx context.Context, VisitorId int) (err error) {
	_, err = dao.Visitor.Ctx(ctx).Where(dao.Visitor.Columns().Id, VisitorId).Delete()
	return
}
