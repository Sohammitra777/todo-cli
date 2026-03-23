package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/Sohammitra777/taskflow/service"
)

func parseUpdateArgs(args []string) (int, string, error) {
	if len(args) < 3 {
		return 0, "", errors.New("invalid arguments")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, "", errors.New("invalid ID format")
	}

	return id, strings.Join(args[2:], " "), nil
}

func printUpdateUsage() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println("  todo update <ID> <text>")

	fmt.Println("\033[1mEXAMPLE\033[0m")
	fmt.Println("  todo update 1 buy gorceries")
}

func handleUpdateError(err error) {
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			fmt.Println("Invalid Task Id")
			return
		} else {
			fmt.Println("Error:", err)
		}
		return
	}
}

func HandleUpdate(s *service.TaskService, args []string) {
	id, desc, err := parseUpdateArgs(args)
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		printUpdateUsage()
		return
	}

	err = s.UpdateTask(desc, id)
	if err != nil {
		handleUpdateError(err)
		return
	}

	fmt.Printf("task updated successfully (ID: %d)\n", id)

}
