package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ayoblt/task-tracker/storage"
)

const DataFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needs more argument")
		os.Exit(1)
	}

	commandName := os.Args[1]

	args := os.Args[2:]

	storage := storage.New(DataFile)

	switch commandName {
	case "add":
		if len(args) != 1 {
			fmt.Println("Add needs a task description to add")
			os.Exit(1)
		}

		err := storage.Add(args[0])
		if err != nil {
			log.Fatal("Task Add failed:", err)
		}

		fmt.Println("Added task")
	case "list":
		tasks, err := storage.Load()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		data, err := json.MarshalIndent(tasks, "", "  ")

		fmt.Println(string(data))
	case "update":
		if len(args) < 2 {
			fmt.Println("requires task ID and description")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid task ID: %s", args[0])
		}

		description := strings.Join(args[1:], " ")

		err = storage.Update(id, description)
		if err != nil {
			fmt.Println("update failed:", err)
			os.Exit(1)
		}

		fmt.Printf("Updated task %d: %s\n", id, description)
	case "delete":
		fmt.Println("Deleted task")
	default:
		fmt.Println("Invalid command.")
	}
}
