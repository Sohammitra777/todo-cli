package types

import "time"

type Status string

const (
	StatusDone       Status = "done"
	StatusInProgress Status = "inProgress"
	StatusNotDone    Status = "NotDone"
)

type Task struct {
	Id        int       `json:"id"`
	Desc      string    `json:"desc"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
