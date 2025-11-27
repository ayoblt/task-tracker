package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayoblt/task-tracker/models"
	"github.com/ayoblt/task-tracker/storage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needs more argument")
		os.Exit(1)
	}

	commandName := os.Args[1]

	args := os.Args[2:]
	switch commandName {
	case "add":
		if len(args) != 1 {
			fmt.Println("Add needs a task description to add")
			os.Exit(1)
		}
		task := models.AddTask(args[0])
		tasks, err := storage.LoadTasks("tasks.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		tasks = append(tasks, task)

		err = storage.SaveTasks("tasks.json", tasks)
		if err != nil {
			fmt.Println("Failed:", err)
		}
		fmt.Println("Added task: ", task)

	case "list":
		fmt.Println("List of tasks")
	case "update":
		fmt.Println("Updated task")
	case "delete":
		fmt.Println("Deleted task")
	default:
		fmt.Println("Invalid command.")
	}
}
