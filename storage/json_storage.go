package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ayoblt/task-tracker/models"
)

func SaveString(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		return err
	}
	return nil
}

func SaveTasks(filename string, tasks []models.Task) error {
	taskJson, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, taskJson, 0644)

	if err != nil {
		return err

	}
	return nil
}

func LoadTasks(filename string) ([]models.Task, error) {
	if !FileExists(filename) {
		return []models.Task{}, nil
	}

	// 1. Read and load file
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []models.Task
	// 2. Decode from JSON
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	// 3. Return tasks struct
	return tasks, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
