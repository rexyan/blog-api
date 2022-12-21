package job

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sJob struct{}

func init() {
	service.RegisterJob(New())
}

func New() *sJob {
	return &sJob{}
}

func (s *sJob) CreateCronJob(ctx context.Context, in *model.CreateCronJobInput) (err error) {
	jobCls := dao.ScheduleJob.Columns()
	_, err = dao.ScheduleJob.Ctx(ctx).Data(g.Map{
		jobCls.BeanName:   in.BeanName,
		jobCls.Cron:       in.Cron,
		jobCls.MethodName: in.MethodName,
		jobCls.Params:     in.Params,
		jobCls.Remark:     in.Remark,
		jobCls.Status:     0,
		jobCls.CreateTime: gtime.Now(),
	}).Insert()
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
	return
}

func (s *sJob) UpdateJobStatus(ctx context.Context, JobId int, JobStatus bool) (err error) {
	_, err = dao.ScheduleJob.Ctx(ctx).Data(g.Map{
		dao.ScheduleJob.Columns().Status: JobStatus,
	}).Where(dao.ScheduleJob.Columns().JobId, JobId).Update()
	return
}

func (s *sJob) UpdateJob(ctx context.Context, in *model.UpdateJobInput) (err error) {
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
	return
}

func (s *sJob) DeleteJob(ctx context.Context, JobId int) (err error) {
	// TODO 停止任务
	_, err = dao.ScheduleJob.Ctx(ctx).Where(dao.ScheduleJob.Columns().JobId, JobId).Delete()
	return
}
