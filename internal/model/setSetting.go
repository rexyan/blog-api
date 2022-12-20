package model

import "blog-api/internal/model/entity"

type UpdateSetSettingInput struct {
	DeleteIds []int                 `json:"deleteIds"`
	Settings  []*entity.SiteSetting `json:"settings"`
}
