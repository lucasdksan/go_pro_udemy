package models

import "time"

type Note struct {
	Id        int
	Title     string
	Content   string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
