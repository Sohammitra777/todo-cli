package cmd

import (
	"errors"
	"fmt"
	"strings"

	"todo.go/service"
)

func parseAddArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("invalid arguments")
	}

	desc := strings.Join(args[1:], " ")
	return desc, nil
}

func printAddUsage() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println(" todo add <text>")
	fmt.Println("\033[1mEXAMPLE\033[0m")
	fmt.Println(" todo add \"Buy Groceries\"")
}

func handleAddError(err error) {
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		return
	}

}

func HandleAdd(s *service.TaskService, args []string) {
	desc, err := parseAddArgs(args)
	if err != nil {
		fmt.Println("\033[3mError:\033[0m", err)
		printAddUsage()
		return
	}

	ID, err := s.AddTask(desc)
	if err != nil {
		handleAddError(err)
		return
	}

	fmt.Printf("Task added successfully (ID : %d)\n", ID)

}
