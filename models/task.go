package models

import "time"

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

func AddTask(description string) Task {
	return Task{
		ID:          2,
		Description: description,
		Status:      TaskStatusTodo,
		CreatedAt:   time.Now(),
	}
}
