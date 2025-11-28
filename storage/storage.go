package storage

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/ayoblt/task-tracker/models"
)

type Storage struct {
	filename string
}

func New(filename string) *Storage {
	return &Storage{filename: filename}
}

func (s *Storage) Load() ([]models.Task, error) {
	if !FileExists(s.filename) {
		return []models.Task{}, nil
	}

	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Storage) Save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Add(description string) error {
	tasks, err := s.Load()

	if err != nil {
		return err
	}
	task := models.CreateTask(description, tasks)

	tasks = append(tasks, task)
	return s.Save(tasks)
}

func (s *Storage) Update(ID int, description string) error {
	found := false
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	var newTasksList []models.Task

	for _, t := range tasks {
		if t.ID == ID {
			found = true
			t.Description = description
			newTasksList = append(newTasksList, t)
		} else {
			newTasksList = append(newTasksList, t)
		}
	}

	if !found {
		return errors.New("Task with ID: " + strconv.Itoa(ID) + " not found")
	}

	return s.Save(newTasksList)
}

func (s *Storage) Delete(ID int) error {
	found := false
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	var newTasksList []models.Task

	for _, t := range tasks {
		if t.ID == ID {
			found = true
			continue
		} else {
			newTasksList = append(newTasksList, t)
		}
	}

	if !found {
		return errors.New("Task with ID: " + strconv.Itoa(ID) + " not found")
	}

	return s.Save(newTasksList)
}

func (s *Storage) MarkInProgress(ID int) error {
	found := false
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	var newTasksList []models.Task

	for _, t := range tasks {
		if t.ID == ID {
			found = true
			t.Status = models.TaskStatusInProgress
			newTasksList = append(newTasksList, t)
		} else {
			newTasksList = append(newTasksList, t)
		}
	}

	if !found {
		return errors.New("Task with ID: " + strconv.Itoa(ID) + " not found")
	}

	return s.Save(newTasksList)
}

func (s *Storage) MarkDone(ID int) error {
	found := false
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	var newTasksList []models.Task

	for _, t := range tasks {
		if t.ID == ID {
			found = true
			t.Status = models.TaskStatusDone
			newTasksList = append(newTasksList, t)
		} else {
			newTasksList = append(newTasksList, t)
		}
	}

	if !found {
		return errors.New("Task with ID: " + strconv.Itoa(ID) + " not found")
	}

	return s.Save(newTasksList)
}

func (s *Storage) MarkTodo(ID int) error {
	found := false
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	var newTasksList []models.Task

	for _, t := range tasks {
		if t.ID == ID {
			found = true
			t.Status = models.TaskStatusTodo
			newTasksList = append(newTasksList, t)
		} else {
			newTasksList = append(newTasksList, t)
		}
	}

	if !found {
		return errors.New("Task with ID: " + strconv.Itoa(ID) + " not found")
	}

	return s.Save(newTasksList)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// func SaveString(filename, content string) error {
// 	err := os.WriteFile(filename, []byte(content), 0644)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func SaveTasks(filename string, tasks []models.Task) error {
// 	taskJson, err := json.Marshal(tasks)

// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(filename, taskJson, 0644)

// 	if err != nil {
// 		return err

// 	}
// 	return nil
// }

// func LoadTasks(filename string) ([]models.Task, error) {
// 	if !FileExists(filename) {
// 		return []models.Task{}, nil
// 	}

// 	// 1. Read and load file
// 	data, err := os.ReadFile(filename)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file: %w", err)
// 	}

// 	var tasks []models.Task
// 	// 2. Decode from JSON
// 	err = json.Unmarshal(data, &tasks)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
// 	}

// 	// 3. Return tasks struct
// 	return tasks, nil
// }
