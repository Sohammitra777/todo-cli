package repo

import (
	"encoding/json"
	"os"

	"github.com/Sohammitra777/taskflow/model"
)

type TaskRepo struct {
	filename string
}

func NewTaskRepo(filename string) *TaskRepo {
	return &TaskRepo{filename: filename}
}

func (r *TaskRepo) CheckAndCreateStorageFile() error {
	if _, err := os.Stat(r.filename); os.IsNotExist(err) {
		f, err := os.Create(r.filename)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}

func (r *TaskRepo) GetTasks() ([]model.Task, error) {
	var tasks []model.Task

	file, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	if len(file) > 0 {
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func (r *TaskRepo) StoreFile(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}
