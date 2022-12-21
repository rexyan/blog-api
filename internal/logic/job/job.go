package job

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"time"
)

type sJob struct{}

func init() {
	service.RegisterJob(New())
}

func New() *sJob {
	return &sJob{}
}

func (s *sJob) CreateCronJob(ctx context.Context, in *model.CreateCronJobInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 新增定时任务，并获取ID
		jobCls := dao.ScheduleJob.Columns()
		jobId, err := dao.ScheduleJob.Ctx(ctx).Data(g.Map{
			jobCls.BeanName:   in.BeanName,
			jobCls.Cron:       in.Cron,
			jobCls.MethodName: in.MethodName,
			jobCls.Params:     in.Params,
			jobCls.Remark:     in.Remark,
			jobCls.Status:     1,
			jobCls.CreateTime: gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			return err
		}

		// 创建定时任务
		jobMethodObject := reflect.ValueOf(&sJob{}).MethodByName(in.MethodName)
		_, err = gcron.Add(ctx, in.Cron, func(ctx context.Context) {
			jobMethodObject.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(gconv.String(jobId)),
				reflect.ValueOf(gconv.String(in.Params)),
			})
		}, gconv.String(jobId))
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *sJob) GetCronJobList(ctx context.Context, page model.PageInput) (res *model.GetCronJobListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetCronJobListOutput{}
	query := dao.ScheduleJob.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetCronJobListOutput{}, err
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

	// 获取已注册的定时任务信息
	g.Log().Info(ctx, "已注册的定时任务:", gcron.Entries())
	return
}

func (s *sJob) UpdateJobStatus(ctx context.Context, JobId int, JobStatus bool) (err error) {
	_, err = dao.ScheduleJob.Ctx(ctx).Data(g.Map{
		dao.ScheduleJob.Columns().Status: JobStatus,
	}).Where(dao.ScheduleJob.Columns().JobId, JobId).Update()

	// 修改定时任务状态
	if JobStatus == true {
		gcron.Start(gconv.String(JobId))
	} else {
		gcron.Stop(gconv.String(JobId))
	}
	return
}

func (s *sJob) UpdateJob(ctx context.Context, in *model.UpdateJobInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 更新定时任务信息
		jobCls := dao.ScheduleJob.Columns()
		_, err = dao.ScheduleJob.Ctx(ctx).Data(g.Map{
			jobCls.BeanName:   in.BeanName,
			jobCls.CreateTime: in.CreateTime,
			jobCls.Cron:       in.Cron,
			jobCls.JobId:      in.JobId,
			jobCls.MethodName: in.MethodName,
			jobCls.Params:     in.Params,
			jobCls.Remark:     in.Remark,
			jobCls.Status:     in.Status,
		}).Where(jobCls.JobId, in.JobId).Update()

		// 删除任务
		gcron.Remove(gconv.String(in.JobId))
		// 重新创建任务
		jobMethodObject := reflect.ValueOf(&sJob{}).MethodByName(in.MethodName)
		_, err = gcron.Add(ctx, in.Cron, func(ctx context.Context) {
			jobMethodObject.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(gconv.String(in.JobId)),
				reflect.ValueOf(gconv.String(in.Params)),
			})
		}, gconv.String(in.JobId))
		if err != nil {
			return err
		}
		// Add 会自动启动任务, 如果是定时任务修改前是关闭状态，则需要关闭任务
		if gconv.Bool(in.Status) == false {
			gcron.Stop(gconv.String(in.JobId))
		}
		return nil
	})
	return
}

func (s *sJob) DeleteJob(ctx context.Context, JobId int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除定时任务记录
		_, err := dao.ScheduleJob.Ctx(ctx).Where(dao.ScheduleJob.Columns().JobId, JobId).Delete()
		if err != nil {
			return err
		}
		// 删除定时任务
		gcron.Remove(gconv.String(JobId))
		return nil
	})
	return
}

func (s *sJob) DryRunJob(ctx context.Context, JobId int) (err error) {
	// 获取 Job 信息
	scheduleJob, err := dao.ScheduleJob.Ctx(ctx).Where(dao.ScheduleJob.Columns().JobId, JobId).One()
	if err != nil {
		return err
	}
	if scheduleJob != nil {
		// 判断是否启用
		if gconv.Int(scheduleJob[dao.ScheduleJob.Columns().Status]) == 0 {
			return gerror.New("状态未启用, 不能被调用!")
		}
		// 获取 job 方法名
		jobMethodName := gconv.String(scheduleJob[dao.ScheduleJob.Columns().MethodName])
		// 获取 job id
		jobId := gconv.String(scheduleJob[dao.ScheduleJob.Columns().JobId])
		// job 参数
		jobParam := gconv.String(scheduleJob[dao.ScheduleJob.Columns().Params])
		// 反射, 根据名称获取函数对象
		jobMethodObject := reflect.ValueOf(&sJob{}).MethodByName(jobMethodName)
		// 触发异步任务, 且延迟 2s 只执行一次
		_, err = gcron.AddOnce(ctx, "@every 3s", func(ctx context.Context) {
			jobMethodObject.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(jobId),
				reflect.ValueOf(jobParam),
			})
		})
		if err != nil {
			return err
		}
		g.Log().Info(ctx, "DryRunJob 触发成功!")
	}
	return
}

func (s *sJob) SyncBlogViewsToDatabase(ctx context.Context, JobId string, Param string) (err error) {
	// 开始时间
	startT := time.Now()
	defer func() {
		status, errMessage := 1, ""
		if err := recover(); err != nil {
			status = 0
			errMessage = gconv.String(err)
		}
		// 记录执行日志
		_ = s.AddJobRecordLog(ctx, &model.AddJobRecordLogInput{
			ScheduleJobLog: entity.ScheduleJobLog{
				JobId:      gconv.Int64(JobId),
				BeanName:   "",
				MethodName: "SyncBlogViewsToDatabase",
				Params:     Param,
				Status:     status,
				Error:      errMessage,
				Times:      gconv.Int(time.Since(startT).Microseconds()),
			},
		})
		g.Log().Info(ctx, "SyncBlogViewsToDatabase 触发成功!")
	}()
	// TODO 实现逻辑

	return
}

func (s *sJob) SyncVisitInfoToDatabase(ctx context.Context, JobId string, Param string) (err error) {
	// 开始时间
	startT := time.Now()
	defer func() {
		status, errMessage := 1, ""
		if err := recover(); err != nil {
			status = 0
			errMessage = gconv.String(err)
		}
		// 记录执行日志
		_ = s.AddJobRecordLog(ctx, &model.AddJobRecordLogInput{
			ScheduleJobLog: entity.ScheduleJobLog{
				JobId:      gconv.Int64(JobId),
				BeanName:   "",
				MethodName: "SyncVisitInfoToDatabase",
				Params:     Param,
				Status:     status,
				Error:      errMessage,
				Times:      gconv.Int(time.Since(startT).Microseconds()),
			},
		})
		g.Log().Info(ctx, "SyncBlogViewsToDatabase 触发成功!")
	}()
	// TODO 实现逻辑

	return
}

func (s *sJob) AddJobRecordLog(ctx context.Context, in *model.AddJobRecordLogInput) (err error) {
	// 记录定时任务日志
	scheduleJobLogCls := dao.ScheduleJobLog.Columns()
	_, err = dao.ScheduleJobLog.Ctx(ctx).Data(
		g.Map{
			scheduleJobLogCls.JobId:      in.JobId,
			scheduleJobLogCls.Status:     in.Status,
			scheduleJobLogCls.Params:     in.Params,
			scheduleJobLogCls.MethodName: in.MethodName,
			scheduleJobLogCls.BeanName:   in.BeanName,
			scheduleJobLogCls.Error:      in.Error,
			scheduleJobLogCls.Times:      in.Times,
			scheduleJobLogCls.CreateTime: gtime.Now(),
		}).Insert()
	return
}

func (s *sJob) InitCronJob(ctx context.Context) (err error) {
	// 查询所有已经注册且启用的任务, 创建定时任务
	allCronJob, err := dao.ScheduleJob.Ctx(ctx).Where(dao.ScheduleJob.Columns().Status, true).All()
	if err != nil {
		return err
	}
	for _, job := range allCronJob {
		// 获取任务 cron，jobId，params，methodName
		cron := gconv.String(job[dao.ScheduleJob.Columns().Cron])
		jobId := gconv.String(job[dao.ScheduleJob.Columns().JobId])
		params := gconv.String(job[dao.ScheduleJob.Columns().JobId])
		methodName := gconv.String(job[dao.ScheduleJob.Columns().MethodName])

		// 创建定时任务
		jobMethodObject := reflect.ValueOf(&sJob{}).MethodByName(methodName)
		_, err = gcron.Add(ctx, cron, func(ctx context.Context) {
			jobMethodObject.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(jobId),
				reflect.ValueOf(params),
			})
		}, gconv.String(jobId))
		if err != nil {
			return err
		}
	}
	return
}
