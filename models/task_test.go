package models

import (
	"fmt"
	"testing"
)

func TestAddTask(t *testing.T) {
	task := AddTask("Learn go with tests")

	if task.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", task.ID)
	}

	if task.Description != "Learn go with tests" {
		t.Errorf("Expected description 'Learn go with tests', got '%s'", task.Description)
	}

	if task.Status != TaskStatusTodo {
		t.Errorf("Expected status to be '%s', got '%s'", TaskStatusTodo, task.Status)
	}

	fmt.Println(task)
}
