package repo

import (
	"encoding/json"
	"os"

	"todo.go/model"
	"todo.go/utils"
)

func CheckAndCreateStorageFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		utils.PrintErr(err)
		defer f.Close()
	}
}

func GetTasks(filename string) ([]model.Task, error) {
	var tasks []model.Task

	file, err := os.ReadFile(filename)
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

func StoreFile(filename string, tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil

}
