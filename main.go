package main

import (
	"fmt"
	"os"
	"todo.go/cmd"
	"todo.go/repo"
	"todo.go/service"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Use -help to determine the working of the cli")
		return
	}

	repo := repo.NewTaskRepo("data/task.json")
	err := repo.CheckAndCreateStorageFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	s := service.NewTaskService(repo)

	switch args[0] {
	case "list":
		cmd.HandleList(s, args)
	case "add":
		cmd.HandleAdd(s, args)
	case "update":
		cmd.HandleUpdate(s, args)
	case "delete":
		cmd.HandleDelete(s, args)
	case "mark":
		cmd.HandleMark(s, args)
	default:
		fmt.Println("No such command exists")
	}
}
