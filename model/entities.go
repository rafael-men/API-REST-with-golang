package model

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"title"`
	Done        bool   `json:"done"`
}
