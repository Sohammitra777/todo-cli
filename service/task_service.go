package service

import (
	"errors"
	"time"

	"todo.go/model"
	"todo.go/repo"
	"todo.go/utils"
)

type TaskService struct {
	repo *repo.TaskRepo
}

func NewTaskService(r *repo.TaskRepo) *TaskService {
	return &TaskService{repo: r}
}

var ErrTaskNotFound = errors.New("task not found")

func (s *TaskService) ListTask() ([]model.Task, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) ListByStatus(status model.Status) ([]model.Task, error) {

	var todoTasks []model.Task
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.Status == status {
			todoTasks = append(todoTasks, t)
		}
	}

	return todoTasks, nil
}

func (s *TaskService) AddTask(desc string) (int, error) {
	task := model.Task{
		Desc:      desc,
		Status:    model.StatusNotDone,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	tasks, err := s.repo.GetTasks()
	if err != nil {
		return 0, err
	}

	task.ID = utils.GetNextId(tasks)
	tasks = append(tasks, task)

	err = s.repo.StoreFile(tasks)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (s *TaskService) UpdateTask(desc string, id int) error {
	tasks, err := s.repo.GetTasks()
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

	err = s.repo.StoreFile(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) MarkTaskStatusById(status model.Status, id int) error {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return ErrTaskNotFound
	}

	err = s.repo.StoreFile(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DeleteTask(id int) (model.Task, error) {

	var task model.Task
	tasks, err := s.repo.GetTasks()
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

	err = s.repo.StoreFile(tasks)
	if err != nil {
		return task, err
	}

	return task, nil
}
