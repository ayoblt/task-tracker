package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ayoblt/task-tracker/models"
	"github.com/ayoblt/task-tracker/storage"
)

const DataFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needs more argument")
		os.Exit(1)
	}

	commandName := os.Args[1]

	if commandName == "help" || commandName == "-h" || commandName == "--help" {
		printUsage()
		return
	}

	args := os.Args[2:]

	storage := storage.New(DataFile)

	switch commandName {
	case "add":
		handleAdd(args, storage)
	case "list":
		handleList(storage, args)
	case "update":
		handleUpdate(args, storage)
	case "delete":
		handleDelete(args, storage)
	case "mark-in-progress":
		handleMarkInProgress(args, storage)
	case "mark-done":
		handleMarkDone(args, storage)
	case "mark-todo":
		handleMarkTodo(args, storage)
	default:
		fmt.Printf("Unknown command: %s\n", commandName)
		printUsage()
		os.Exit(1)
	}
}

func handleAdd(args []string, storage *storage.Storage) {
	if len(args) < 1 {
		fmt.Println("Add needs a task description to add")
		os.Exit(1)
	}

	description := strings.Join(args, " ")

	err := storage.Add(description)
	if err != nil {
		log.Fatal("Task Add failed:", err)
	}

	fmt.Println("Added task")
}

func handleList(storage *storage.Storage, args []string) {
	tasks, err := storage.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Filter by status if provided
	var filter models.TaskStatus
	if len(args) > 0 {
		filter = models.TaskStatus(args[0])
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	count := 0
	for _, task := range tasks {
		// Skip if filter doesn't match
		if filter != "" && task.Status != filter {
			continue
		}

		status := formatStatus(task.Status)
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Description)
		count++
	}

	if count == 0 {
		fmt.Printf("No tasks with status '%s'\n", filter)
	}
}
func formatStatus(status models.TaskStatus) string {
	switch status {
	case models.TaskStatusTodo:
		return "[ ]"
	case models.TaskStatusInProgress:
		return "[~]"
	case models.TaskStatusDone:
		return "[✓]"
	default:
		return "[?]"
	}
}

func handleUpdate(args []string, storage *storage.Storage) {
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
}

func handleDelete(args []string, storage *storage.Storage) {
	if len(args) < 1 {
		fmt.Println("requires task ID")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("invalid task ID: %s", args[0])
	}

	err = storage.Delete(id)
	if err != nil {
		fmt.Println("delete failed:", err)
		os.Exit(1)
	}

	fmt.Printf("Deleted task %d\n", id)
}

func handleMarkInProgress(args []string, storage *storage.Storage) {
	if len(args) < 1 {
		fmt.Println("requires task ID")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("invalid task ID: %s", args[0])
	}

	err = storage.MarkInProgress(id)
	if err != nil {
		fmt.Println("Mark in-progress failed:", err)
		os.Exit(1)
	}

	fmt.Printf("task %d set to in-progress\n", id)

}

func handleMarkDone(args []string, storage *storage.Storage) {
	if len(args) < 1 {
		fmt.Println("requires task ID")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("invalid task ID: %s", args[0])
	}

	err = storage.MarkDone(id)
	if err != nil {
		fmt.Println("Mark done failed:", err)
		os.Exit(1)
	}

	fmt.Printf("task %d set to done\n", id)

}

func handleMarkTodo(args []string, storage *storage.Storage) {
	if len(args) < 1 {
		fmt.Println("requires task ID")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("invalid task ID: %s", args[0])
	}

	err = storage.MarkTodo(id)
	if err != nil {
		fmt.Println("Mark todo failed:", err)
		os.Exit(1)
	}

	fmt.Printf("task %d set to todo\n", id)

}
func printUsage() {
	usage := `
Task Tracker - A simple CLI task manager

Usage:
  task-tracker add <description>           Add a new task
  task-tracker list [status]               List all tasks or filter by status
  task-tracker update <id> <description>   Update a task
  task-tracker delete <id>                 Delete a task
  task-tracker mark-done <id>              Mark task as done
  task-tracker mark-in-progress <id>       Mark task as in progress

Examples:
  task-tracker add "Buy groceries"
  task-tracker list
  task-tracker list done
  task-tracker update 1 "Buy groceries and cook dinner"
  task-tracker mark-done 1
`
	fmt.Println(usage)
}
