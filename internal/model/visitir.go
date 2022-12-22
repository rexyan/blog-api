package model

import "blog-api/internal/model/entity"

type TodayPvUvOutput struct {
	Pv int
	Uv int
}

type LatestMonthPvUvOutput struct {
	Date string
	Pv   int
	Uv   int
}

type DashBoardMonthPvUv struct {
	Date []string
	Pv   []int
	Uv   []int
}

type CityVisitorOutput struct {
	City string
	Uv   int
}

type GetVisitListOutput struct {
	CommonPageHelper
	List []*entity.Visitor
}
