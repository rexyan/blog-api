package model

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
