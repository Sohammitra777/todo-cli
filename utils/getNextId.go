package utils

import (
	"todo.go/model"
)

func GetNextId(tasks []model.Task) int {

	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
