package log

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLog struct{}

func init() {
	service.RegisterLog(New())
}

func New() *sLog {
	return &sLog{}
}

func (s *sLog) GetJobLogList(ctx context.Context, page model.PageInput, startTime *gtime.Time, endTime *gtime.Time) (res *model.GetJobLogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetJobLogListOutput{}
	query := dao.ScheduleJobLog.Ctx(ctx).OrderDesc(dao.ScheduleJobLog.Columns().LogId)
	// 按照时间过滤
	if startTime != nil && endTime != nil {
		query = query.WhereBetween(dao.ScheduleJobLog.Columns().CreateTime, startTime, endTime)
	}

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetJobLogListOutput{}, err
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

func (s *sLog) DeleteJobLog(ctx context.Context, LogId int) (err error) {
	_, err = dao.ScheduleJobLog.Ctx(ctx).Where(dao.ScheduleJobLog.Columns().LogId, LogId).Delete()
	return
}

func (s *sLog) GetLoginLogList(ctx context.Context, page model.PageInput, startTime *gtime.Time, endTime *gtime.Time) (res *model.GetLoginLogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetLoginLogListOutput{}
	query := dao.LoginLog.Ctx(ctx).OrderDesc(dao.LoginLog.Columns().Id)
	// 按照时间过滤
	if startTime != nil && endTime != nil {
		query = query.WhereBetween(dao.LoginLog.Columns().CreateTime, startTime, endTime)
	}
	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetLoginLogListOutput{}, err
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

func (s *sLog) DeleteLoginLog(ctx context.Context, LogId int) (err error) {
	_, err = dao.LoginLog.Ctx(ctx).Where(dao.LoginLog.Columns().Id, LogId).Delete()
	return
}

func (s *sLog) GetOperationLogList(ctx context.Context, page model.PageInput, startTime *gtime.Time, endTime *gtime.Time) (res *model.GetOperationLogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetOperationLogListOutput{}
	query := dao.OperationLog.Ctx(ctx).OrderDesc(dao.OperationLog.Columns().Id)
	// 按照时间过滤
	if startTime != nil && endTime != nil {
		query = query.WhereBetween(dao.OperationLog.Columns().CreateTime, startTime, endTime)
	}
	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetOperationLogListOutput{}, err
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

func (s *sLog) DeleteOperationLog(ctx context.Context, LogId int) (err error) {
	_, err = dao.OperationLog.Ctx(ctx).Where(dao.OperationLog.Columns().Id, LogId).Delete()
	return
}

func (s *sLog) GetExceptionLogList(ctx context.Context, page model.PageInput, startTime *gtime.Time, endTime *gtime.Time) (res *model.GetExceptionLogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetExceptionLogListOutput{}
	query := dao.ExceptionLog.Ctx(ctx).OrderDesc(dao.ExceptionLog.Columns().Id)
	// 按照时间过滤
	if startTime != nil && endTime != nil {
		query = query.WhereBetween(dao.ExceptionLog.Columns().CreateTime, startTime, endTime)
	}
	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetExceptionLogListOutput{}, err
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

func (s *sLog) DeleteExceptionLog(ctx context.Context, LogId int) (err error) {
	_, err = dao.ExceptionLog.Ctx(ctx).Where(dao.ExceptionLog.Columns().Id, LogId).Delete()
	return
}

func (s *sLog) GetVisitLogList(ctx context.Context, page model.PageInput, startTime *gtime.Time, endTime *gtime.Time, VisitorId string) (res *model.GetVisitLogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetVisitLogListOutput{}
	query := dao.VisitLog.Ctx(ctx).OrderDesc(dao.VisitLog.Columns().Id)
	// 按照时间过滤
	if startTime != nil && endTime != nil {
		query = query.WhereBetween(dao.VisitLog.Columns().CreateTime, startTime, endTime)
	}
	// 按照访客标识过滤
	if len(VisitorId) > 0 {
		query = query.Where(dao.VisitLog.Columns().Uuid, VisitorId)
	}
	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetVisitLogListOutput{}, err
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

func (s *sLog) DeleteVisitLog(ctx context.Context, LogId int) (err error) {
	_, err = dao.VisitLog.Ctx(ctx).Where(dao.VisitLog.Columns().Id, LogId).Delete()
	return
}
