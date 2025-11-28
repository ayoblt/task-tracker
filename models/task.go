package models

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateTask(description string, tasks []Task) Task {
	taskID := nextID(tasks)

	return Task{
		ID:          taskID,
		Description: description,
		Status:      TaskStatusTodo,
		CreatedAt:   time.Now(),
	}
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1
}
