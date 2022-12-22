package model

import "blog-api/internal/model/entity"

type GetJobLogListOutput struct {
	CommonPageHelper
	List []*entity.ScheduleJobLog
}

type GetLoginLogListOutput struct {
	CommonPageHelper
	List []*entity.LoginLog
}

type GetOperationLogListOutput struct {
	CommonPageHelper
	List []*entity.OperationLog
}

type GetExceptionLogListOutput struct {
	CommonPageHelper
	List []*entity.ExceptionLog
}

type GetVisitLogListOutput struct {
	CommonPageHelper
	List []*entity.VisitLog
}
