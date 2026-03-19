package main

import (
	"time"
	"todo.go/types"
	"todo.go/utils"
)

func main() {
	filename := "task.json"
	utils.CheckAndCreateStorageFile(filename)

	task := types.Task{
		Id:        1,
		Desc:      "Learn Go",
		Status:    types.StatusNotDone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var tasks []types.Task
	tasks = append(tasks, task)
	utils.SaveToJson(filename, tasks)

}
