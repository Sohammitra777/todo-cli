package service

import (
	"time"

	"todo.go/model"
	"todo.go/repo"
	"todo.go/utils"
)

func AddTask(filename string, desc string) (int, error) {
	task := model.Task{
		Desc:      desc,
		Status:    model.StatusNotDone,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	tasks, err := repo.GetTasks(filename)
	if err != nil {
		return 0, err
	}

	task.ID = utils.GetNextId(tasks)
	tasks = append(tasks, task)

	err = repo.StoreFile(filename, tasks)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}
