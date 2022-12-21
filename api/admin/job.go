package admin

import "github.com/gogf/gf/v2/frame/g"

type CreateCronJobReq struct {
	g.Meta        `path:"/job" tags:"定时任务" method:"post" summary:"新增定时任务"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	BeanName      string `json:"beanName"`
	MethodName    string `json:"methodName"`
	Params        string `json:"params"`
	Cron          string `json:"cron"`
	Remark        string `json:"remark"`
}

type CreateCronJobRes struct {
}

type GetCronJobListReq struct {
	g.Meta        `path:"/jobs" tags:"定时任务" method:"get" summary:"定时任务列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetCronJobListRes struct {
	CommonPageHelper
	List []Job `json:"list"`
}

type Job struct {
	JobID      int    `json:"jobId"`
	BeanName   string `json:"beanName"`
	MethodName string `json:"methodName"`
	Params     string `json:"params"`
	Cron       string `json:"cron"`
	Status     bool   `json:"status"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}

type UpdateJobStatusReq struct {
	g.Meta        `path:"/job/status" tags:"定时任务" method:"put" summary:"更新定时任务状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	JobID         int    `json:"jobId"`
	Status        bool   `json:"status"`
}

type UpdateJobStatusRes struct {
}

type UpdateJobReq struct {
	g.Meta        `path:"/job" tags:"定时任务" method:"put" summary:"更新定时任务"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Job
}

type UpdateJobRes struct {
}

type DeleteJobReq struct {
	g.Meta        `path:"/job" tags:"定时任务" method:"delete" summary:"删除定时任务"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	JobID         int    `json:"jobId"`
}

type DeleteJobRes struct {
}
