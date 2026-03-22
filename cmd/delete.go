package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"todo.go/model"
	"todo.go/service"
	"todo.go/utils"
)

func parseDeleteArgs(args []string) (int, error) {
	if len(args) != 2 {
		return 0, errors.New("invalid arguments")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, errors.New("invalid ID format")

	}
	return id, nil
}

func printDeleteUsage() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println("  todo delete <ID>")
	fmt.Println("\033[1mEXAMPLE\033[0m")
	fmt.Println("  todo delete 1")
}

func handleDeleteError(err error) {
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			fmt.Println("Task not found")
		} else {
			fmt.Println("Error:", err)
		}
	}
}

func printDeletedTask(task model.Task) {
	fmt.Println("Task delete successfully")
	fmt.Println("------deleted task------")
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("%s\n", utils.CapitalizeWords(task.Desc))
	fmt.Printf("Status : %s\n", task.Status)
}

func HandleDelete(s *service.TaskService, args []string) {

	id, err := parseDeleteArgs(args)
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		printDeleteUsage()
		return
	}

	task, err := s.DeleteTask(id)
	if err != nil {
		handleDeleteError(err)
		return
	}

	printDeletedTask(task)
}
