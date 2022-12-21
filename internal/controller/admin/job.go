package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var CronJob = cCronJob{}

type cCronJob struct{}

func (c *cCronJob) CreateJob(ctx context.Context, req *admin.CreateCronJobReq) (res *admin.CreateCronJobRes, err error) {
	err = service.Job().CreateCronJob(ctx, &model.CreateCronJobInput{
		ScheduleJob: entity.ScheduleJob{
			BeanName:   req.BeanName,
			MethodName: req.MethodName,
			Params:     req.Params,
			Cron:       req.Cron,
			Remark:     req.Remark,
		},
	})
	return
}

func (c *cCronJob) GetJobList(ctx context.Context, req *admin.GetCronJobListReq) (res *admin.GetCronJobListRes, err error) {
	jobList, err := service.Job().GetCronJobList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	err = gconv.Scan(jobList, &res)
	return
}

func (c *cCronJob) UpdateJobStatus(ctx context.Context, req *admin.UpdateJobStatusReq) (res *admin.UpdateJobStatusRes, err error) {
	err = service.Job().UpdateJobStatus(ctx, req.JobID, req.Status)
	return
}

func (c *cCronJob) UpdateJob(ctx context.Context, req *admin.UpdateJobReq) (res *admin.UpdateJobRes, err error) {
	err = service.Job().UpdateJob(ctx, &model.UpdateJobInput{
		ScheduleJob: entity.ScheduleJob{
			JobId:      gconv.Int64(req.JobID),
			BeanName:   req.BeanName,
			MethodName: req.MethodName,
			Params:     req.Params,
			Cron:       req.Cron,
			Status:     gconv.Int(req.Status),
			Remark:     req.Remark,
			CreateTime: gtime.New(req.CreateTime),
		},
	})
	return
}

func (c *cCronJob) DeleteJob(ctx context.Context, req *admin.DeleteJobReq) (res *admin.DeleteJobRes, err error) {
	err = service.Job().DeleteJob(ctx, req.JobID)
	return
}

func (c *cCronJob) DryRunJob(ctx context.Context, req *admin.DryRunJobReq) (res *admin.DryRunJobRes, err error) {
	err = service.Job().DryRunJob(ctx, req.JobID)
	return
}
