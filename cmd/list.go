package cmd

import (
	"fmt"
	"github.com/Sohammitra777/taskflow/model"
	"github.com/Sohammitra777/taskflow/service"
	"github.com/Sohammitra777/taskflow/utils"
)

func printListUsage() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println("  todo list <subcommand>")

	fmt.Println("\033[1mSUBCOMMANDS\033[0m")
	fmt.Println("  todo")
	fmt.Println("  in-progress")
	fmt.Println("  done")
}

func printTaskList(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("Task list empty")
		return
	}

	fmt.Printf("Total number of task : %d\n", len(tasks))
	fmt.Println("=======")
	fmt.Println("o---List start---o")
	for i, task := range tasks {
		if i != 0 {
			fmt.Println("------")
		}
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("%s\n", utils.CapitalizeWords(task.Desc))
		fmt.Printf("Status : %s\n", task.Status)

	}
	fmt.Println("x-------End-------x")

}

func handleListMethod(s *service.TaskService, subcommand string) {
	var listFilters = map[string]model.Status{
		"todo":        model.StatusNotDone,
		"in-progress": model.StatusInProgress,
		"done":        model.StatusDone,
	}
	status, ok := listFilters[subcommand]
	if ok {
		tasks, err := s.ListByStatus(status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printTaskList(tasks)

	} else {
		fmt.Print("\033[3mError:\033[0m")
		fmt.Print(" invalid subcommand\n")
		printListUsage()
	}
}

func HandleList(s *service.TaskService, args []string) {

	if len(args) == 1 {
		tasks, err := s.ListTask()
		if err != nil {
			fmt.Println("\033[3mError:\033[0m", "invalid arguments")
			return
		}

		printTaskList(tasks)
		return
	}

	if len(args) == 2 {
		handleListMethod(s, args[1])
		return
	}

	fmt.Println("\033[3mError:\033[0", "invalid arguments")
	printListUsage()
}
