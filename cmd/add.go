package cmd

import (
	"fmt"
	"strings"
	"todo.go/service"
	"todo.go/utils"
)

func HandleAdd(filename string, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: todo add <value>")
		fmt.Println("Example Usage : todo add \"Buy Groceries\"")
		return
	}

	desc := strings.Join(args[1:], " ")
	ID, err := service.AddTask(filename, desc)
	utils.PrintErr(err)

	fmt.Printf("Task added successfully (ID : %d)\n", ID)
}
