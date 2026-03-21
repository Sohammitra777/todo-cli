package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"todo.go/service"
	"todo.go/utils"
)

func HandleUpdate(filename string, args []string) {
	if len(args) < 3 {
		fmt.Println("No such command exist")
		return
	}

	id, err := strconv.Atoi(args[1])
	utils.PrintErr(err)
	desc := strings.Join(args[2:], " ")

	err = service.UpdateTask(filename, desc, id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			fmt.Println("Invalid Task Id")
			return
		} else {
			utils.PrintErr(err)
		}
		return
	}

	fmt.Printf("task updated successfully (ID: %d)\n", id)

}
