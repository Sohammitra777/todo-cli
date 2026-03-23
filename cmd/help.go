package cmd

import "fmt"

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func dim(s string) string {
	return "\033[2m" + s + "\033[0m"
}

func HandleHelp() {
	fmt.Println(bold("Taskflow"))
	fmt.Println(dim("A simple command-line todo manager"))
	fmt.Println(dim("Made by Soham Mitra"))
	fmt.Println(dim("GitHub: https://github.com/Sohammitra777"))
	fmt.Println(dim("Repo:   https://github.com/Sohammitra777/todo-cli"))
	fmt.Println()

	fmt.Println(bold("Usage"))

	fmt.Println(dim("list") + ":")
	fmt.Println("  todo list <subcommand>")

	fmt.Println(dim("subcommands") + ":")
	fmt.Println("  todo")
	fmt.Println("  in-progress")
	fmt.Println("  done")

	fmt.Println(dim("add") + ":")
	fmt.Println("  todo add <text>")

	fmt.Println(dim("update") + ":")
	fmt.Println("  todo update <ID> <text>")

	fmt.Println(dim("delete") + ":")
	fmt.Println("  todo delete <ID>")

	fmt.Println(dim("mark") + ":")
	fmt.Println("  todo mark <subcommand>")

	fmt.Println(dim("subcommands") + ":")
	fmt.Println("  todo")
	fmt.Println("  in-progress")
	fmt.Println("  done")
}