package list

import (
	"todo.go/service"
	"todo.go/utils"
)

func HandleListInProgress(filename string) {
	tasks, err := service.TaskListInProgress(filename)
	utils.PrintErr(err)
	
	utils.PrintTaskList(tasks)
}