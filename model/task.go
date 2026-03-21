package model

import "time"

type Status string

const (
	StatusDone       Status = "done"
	StatusInProgress Status = "In Progress"
	StatusNotDone    Status = "Not Done"
)

type Task struct {
	ID        int       `json:"id"`
	Desc      string    `json:"desc"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
