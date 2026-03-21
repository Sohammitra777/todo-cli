package cmd

import (
	"fmt"

	"todo.go/cmd/list"
	"todo.go/service"
	"todo.go/utils"
)

func HandleList(filename string, args []string) {

	if len(args) < 2 {
		tasks, err := service.ListTask(filename)
		utils.PrintErr(err)

		utils.PrintTaskList(tasks)
		return
	}

	switch args[1] {
	case "todo":
		list.HandleListTodo(filename)
	case "in-progress":
		list.HandleListInProgress(filename)
	case "done":
		list.HandleListDone(filename)
	default:
		fmt.Println("No such subcommand exist")

	}

}
