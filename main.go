package main

import (
	"fmt"
	"os"

	"todo.go/cmd"
	"todo.go/repo"
)

func main() {
	filename := "task.json"
	repo.CheckAndCreateStorageFile(filename)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Use -help to determine the working of the cli")
		return
	}

	switch args[0] {
	case "list":
		cmd.HandleList(filename, args)
	case "add":
		cmd.HandleAdd(filename, args)

	case "update":
		cmd.HandleUpdate(filename, args)

	case "delete":
		cmd.HandleDelete(filename, args)
	default:
		fmt.Println("No such command exists")

	}
}
