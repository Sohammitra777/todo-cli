package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"todo.go/service"
	"todo.go/utils"
)

func HandleDelete(filename string, args []string) {

	if len(args) > 2 {
		fmt.Println("No such command exist")
		return
	}

	if len(args) < 2 {
		fmt.Println("Invalid format")
		fmt.Println("\033[1mUSAGE\033[0m")
		fmt.Println(" todo delete <ID>")
		fmt.Println("\033[1mEXAMPLE USAGE\033[0m")
		fmt.Println(" todo delete 1")
		return
	}

	id, err := strconv.Atoi(args[1])
	utils.PrintErr(err)

	task, err := service.DeleteTask(filename, id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			fmt.Println("Task not found")
		} else {
			utils.PrintErr(err)
		}
		return
	}

	utils.PrintDeletedTask(task)
}
