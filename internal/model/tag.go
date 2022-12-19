package model

type Series struct {
	ID    int
	Name  string
	Value int
}

type DashboardTagOutput struct {
	Legend []string
	Series []Series
}

type CreateTagInput struct {
	Color   string
	TagName string
}

type CreateTagOutput struct {
}

type GetTagListOutput struct {
	CommonPageHelper
	List []struct {
		ID int
		CreateTagInput
	}
}

type UpdateTagInput struct {
	ID int
	CreateTagInput
}
