package model

type TodayPvUvOutput struct {
	Pv int
	Uv int
}

type LatestMonthPvUvOutput struct {
	Pv []int
	Uv []int
}

type CityVisitorOutput struct {
	City string
	Uv   int
}
