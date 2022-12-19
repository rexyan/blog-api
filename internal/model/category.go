package model

import "blog-api/internal/model/entity"

type DashboardCategoryOutput struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}

type GteCategoryListOutPut struct {
	CommonPageHelper
	List []entity.Category
}
