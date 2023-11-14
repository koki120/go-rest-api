package entity

import "time"

type Memo struct {
	MemoID    string
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type MemoCrate struct {
	Title string
	Body  string
}
