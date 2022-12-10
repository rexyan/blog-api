package model

type DashboardCategoryOutput struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
