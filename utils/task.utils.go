package utils

import (
	"fmt"
	"strings"
	"unicode"

	"todo.go/model"
)

func GetNextId(tasks []model.Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

func PrintErr(err error) {
	if err != nil {
		panic(err)
	}
}

func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
		words[i] = string(r)
	}
	return strings.Join(words, " ")
}

func PrintTaskList(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("Task list empty")
		return
	}

	fmt.Printf("Total number of task : %d\n", len(tasks))
	fmt.Println("=======")
	fmt.Println("o---List start---o")
	for i, task := range tasks {
		if i != 0 {
			fmt.Println("------")
		}
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("%s\n", capitalizeWords(task.Desc))
		fmt.Printf("Status : %s\n", task.Status)

	}
	fmt.Println("x-------End-------x")

}

func PrintDeletedTask(task model.Task) {
	fmt.Println("Task delete successfully")
	fmt.Println("------deleted task------")
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("%s\n", capitalizeWords(task.Desc))
	fmt.Printf("Status : %s\n", task.Status)
}
