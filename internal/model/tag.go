package model

type Series struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type DashboardTagOutput struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
