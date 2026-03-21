package cmd

import (
	"fmt"

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

	fmt.Println("No such command exist")

}
