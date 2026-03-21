package main

import (
	"fmt"
	"os"

	"todo.go/cmd"
	"todo.go/repo"
)

func main() {
	filename := "task.json"
	commands := [4]string{"list", "add", "update", "delete"}
	repo.CheckAndCreateStorageFile(filename)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Use -help to determine the working of the cli")
		return
	}

	command := args[0]
	switch command {
	case commands[0]:
		cmd.HandleList(filename, args)
	case commands[1]:
		cmd.HandleAdd(filename, args)

	case commands[2]:
		cmd.HandleUpdate(filename, args)

	case commands[3]:
		cmd.HandleDelete(filename, args)
	default:
		fmt.Println("No such command exists")

	}
}
