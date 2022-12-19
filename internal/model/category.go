package model

import "blog-api/internal/model/entity"

type DashboardCategoryOutput struct {
	Legend []string
	Series []Series
}

type GteCategoryListOutPut struct {
	CommonPageHelper
	List []entity.Category
}
