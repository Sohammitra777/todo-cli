package main

import (
	"fmt"
	"os"

	"github.com/Sohammitra777/taskflow/cmd"
	"github.com/Sohammitra777/taskflow/repo"
	"github.com/Sohammitra777/taskflow/service"
)

func printCommands() {
	fmt.Println("\033[1mUSAGE\033[0m")
	fmt.Println("  todo <command>")

	fmt.Println("\033[1mCOMMANDS\033[0m")
	fmt.Println("  help")
	fmt.Println("  list")
	fmt.Println("  add")
	fmt.Println("  mark")
	fmt.Println("  update")
	fmt.Println("  delete")

}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("\033[3mError: \033[0m")
		fmt.Print("invalid command\n")
		printCommands()
		return
	}

	repo := repo.NewTaskRepo("data/task.json")
	err := repo.CheckAndCreateStorageFile()
	if err != nil {
		fmt.Println("\033[3mError: \033[0m", err)
		return
	}
	s := service.NewTaskService(repo)

	switch args[0] {
	case "help":
		cmd.HandleHelp()
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
		{
			fmt.Print("\033[3mError: \033[0m")
			fmt.Print("invalid command\n")
			printCommands()

		}
	}
}
