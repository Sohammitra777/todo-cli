package list

import (
	"todo.go/service"
	"todo.go/utils"
)

func HandleListTodo(filename string) {
	tasks, err := service.TaskListTodo(filename)
	utils.PrintErr(err)

	utils.PrintTaskList(tasks)
}
