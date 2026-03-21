package list

import (
	"todo.go/service"
	"todo.go/utils"
)

func HandleListDone(filename string) {
	tasks, err := service.TaskListDone(filename)
	utils.PrintErr(err)

	utils.PrintTaskList(tasks)
}
