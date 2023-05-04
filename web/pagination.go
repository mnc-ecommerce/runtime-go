package web

type Pagination struct {
	Page  Page `json:"page"`
	Total int  `json:"total"`
	Limit int  `json:"limit"`
}

type Page struct {
	Current  int `json:"current"`
	Previous int `json:"previous"`
	Next     int `json:"next"`
	Total    int `json:"total"`
}
