package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

var Log = cLog{}

type cLog struct{}

func (c *cLog) GetJobLogList(ctx context.Context, req *admin.GetJobLogListReq) (res *admin.GetJobLogListRes, err error) {
	var startTime *gtime.Time
	var endTime *gtime.Time

	if len(req.Date) > 0 {
		startTime = gconv.GTime(strings.Split(req.Date, ",")[0])
		endTime = gconv.GTime(strings.Split(req.Date, ",")[1])
	}
	jobLogList, err := service.Log().GetJobLogList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, startTime, endTime)
	err = gconv.Scan(jobLogList, &res)
	return
}

func (c *cLog) DeleteJobLog(ctx context.Context, req *admin.DeleteJobLogReq) (res *admin.DeleteJobLogRes, err error) {
	err = service.Log().DeleteJobLog(ctx, req.LogId)
	return
}

func (c *cLog) GetLoginLogList(ctx context.Context, req *admin.GetLoginLogListReq) (res *admin.GetLoginLogListRes, err error) {
	var startTime *gtime.Time
	var endTime *gtime.Time

	if len(req.Date) > 0 {
		startTime = gconv.GTime(strings.Split(req.Date, ",")[0])
		endTime = gconv.GTime(strings.Split(req.Date, ",")[1])
	}
	loginLogList, err := service.Log().GetLoginLogList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, startTime, endTime)
	err = gconv.Scan(loginLogList, &res)
	return
}

func (c *cLog) DeleteLoginLog(ctx context.Context, req *admin.DeleteLoginLogReq) (res *admin.DeleteLoginLogRes, err error) {
	err = service.Log().DeleteLoginLog(ctx, req.LogId)
	return
}

func (c *cLog) GetOperationLogList(ctx context.Context, req *admin.GetOperationLogListReq) (res *admin.GetOperationLogListRes, err error) {
	var startTime *gtime.Time
	var endTime *gtime.Time

	if len(req.Date) > 0 {
		startTime = gconv.GTime(strings.Split(req.Date, ",")[0])
		endTime = gconv.GTime(strings.Split(req.Date, ",")[1])
	}
	operationLogList, err := service.Log().GetOperationLogList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, startTime, endTime)
	err = gconv.Scan(operationLogList, &res)
	return
}

func (c *cLog) DeleteOperationLog(ctx context.Context, req *admin.DeleteOperationLogReq) (res *admin.DeleteOperationLogRes, err error) {
	err = service.Log().DeleteOperationLog(ctx, req.LogId)
	return
}

func (c *cLog) GetExceptionLogList(ctx context.Context, req *admin.GetExceptionLogListReq) (res *admin.GetExceptionLogListRes, err error) {
	var startTime *gtime.Time
	var endTime *gtime.Time

	if len(req.Date) > 0 {
		startTime = gconv.GTime(strings.Split(req.Date, ",")[0])
		endTime = gconv.GTime(strings.Split(req.Date, ",")[1])
	}
	exceptionLogList, err := service.Log().GetExceptionLogList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, startTime, endTime)
	err = gconv.Scan(exceptionLogList, &res)
	return
}

func (c *cLog) DeleteExceptionLog(ctx context.Context, req *admin.DeleteExceptionLogReq) (res *admin.DeleteExceptionLogRes, err error) {
	err = service.Log().DeleteExceptionLog(ctx, req.LogId)
	return
}

func (c *cLog) GetVisitLogList(ctx context.Context, req *admin.GetVisitLogListReq) (res *admin.GetVisitLogListRes, err error) {
	var startTime *gtime.Time
	var endTime *gtime.Time

	if len(req.Date) > 0 {
		startTime = gconv.GTime(strings.Split(req.Date, ",")[0])
		endTime = gconv.GTime(strings.Split(req.Date, ",")[1])
	}
	visitLogList, err := service.Log().GetVisitLogList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, startTime, endTime, req.VisitorId)
	err = gconv.Scan(visitLogList, &res)
	return
}

func (c *cLog) DeleteVisitLog(ctx context.Context, req *admin.DeleteVisitLogReq) (res *admin.DeleteVisitLogRes, err error) {
	err = service.Log().DeleteVisitLog(ctx, req.LogId)
	return
}
