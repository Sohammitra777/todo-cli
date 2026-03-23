package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Sohammitra777/taskflow/model"
	"github.com/Sohammitra777/taskflow/service"
)

func printMarkUsage() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println("  todo mark <ID> <status>")

	fmt.Println("\033[1mSTATUS\033[0m")
	fmt.Println("  todo")
	fmt.Println("  in-progress")
	fmt.Println("  done")
}

func parseMarkArgs(args []string) (int, error) {
	if len(args) == 2 {
		return 0, errors.New("invalid status")
	}

	if len(args) != 3 {
		return 0, errors.New("invlid arguments")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, errors.New("Invalid ID format")
	}

	return id, nil
}

func HandleMark(s *service.TaskService, args []string) {

	id, err := parseMarkArgs(args)
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		printMarkUsage()
		return
	}

	markFilters := map[string]model.Status{
		"todo":        model.StatusNotDone,
		"in-progress": model.StatusInProgress,
		"done":        model.StatusDone,
	}
	status, ok := markFilters[args[2]]
	if !ok {
		fmt.Print("\033[3mError:\033[0m")
		fmt.Println("invalid status")
		printMarkUsage()
		return
	}

	err = s.MarkTaskStatusById(status, id)
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		return
	}

	fmt.Println("Status updated successfully")

}
