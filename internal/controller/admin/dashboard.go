package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var Dashboard = cDashboard{}

type cDashboard struct{}

func (c *cDashboard) Index(ctx context.Context, req *admin.DashboardReq) (res *admin.DashboardRes, err error) {
	res = &admin.DashboardRes{}
	pvUv, err := service.Visitor().GetTodayPvUv(ctx)
	blogCount := service.Blog().GetBlogCount(ctx)
	commentCount := service.Comment().GetCommentCount(ctx)
	cityVisitor, err := service.Visitor().GetCityVisitor(ctx)
	tagStatistics, err := service.Tag().DashboardTagStatistics(ctx)
	categoryStatistics, err := service.Category().DashboardCategoryStatistics(ctx)
	visitStatistics, err := service.Visitor().DashBoardMonthPvUvStatistics(ctx)

	if pvUv != nil {
		res.Pv = pvUv.Pv
		res.Uv = pvUv.Uv
	}

	res.BlogCount = gconv.Int(blogCount)
	res.CommentCount = gconv.Int(commentCount)
	err = gconv.Scan(cityVisitor, &res.CityVisitor)
	err = gconv.Scan(tagStatistics, &res.Tag)
	err = gconv.Scan(categoryStatistics, &res.Category)
	err = gconv.Scan(visitStatistics, &res.VisitRecord)

	if err != nil {
		return nil, err
	}
	return
}
