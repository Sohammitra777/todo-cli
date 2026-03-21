package service

import (
	"errors"
	"time"

	"todo.go/model"
	"todo.go/repo"
	"todo.go/utils"
)

var ErrTaskNotFound = errors.New("task not found")

func ListTask(filename string) ([]model.Task, error) {
	tasks, err := repo.GetTasks(filename)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

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

func UpdateTask(filename, desc string, id int) error {
	tasks, err := repo.GetTasks(filename)
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Desc = desc
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return ErrTaskNotFound
	}

	err = repo.StoreFile(filename, tasks)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(filename string, id int) (model.Task, error) {

	var task model.Task
	tasks, err := repo.GetTasks(filename)
	if err != nil {
		return task, err
	}

	idx := -1
	for i, t := range tasks {
		if t.ID == id {
			idx = i
		}
	}

	if idx == -1 {
		return task, ErrTaskNotFound
	}

	task = tasks[idx]
	tasks = append(tasks[:idx], tasks[idx+1:]...)

	err = repo.StoreFile(filename, tasks)
	if err != nil {
		return task, err
	}

	return task, nil
}
