package models

import "time"

type Todo struct {
	Title      string    `json:"title"`
	IsDone     bool      `json:"is_done"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
}
