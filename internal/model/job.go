package model

import "blog-api/internal/model/entity"

type CreateCronJobInput struct {
	entity.ScheduleJob
}

type GetCronJobListOutput struct {
	CommonPageHelper
	List []*entity.ScheduleJob
}

type UpdateJobInput struct {
	entity.ScheduleJob
}
