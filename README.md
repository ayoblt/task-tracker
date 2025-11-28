# Task Tracker CLI

A simple command-line task management application built with Go to practice file I/O, testing, and CLI development.

## Why This Project?

This project was built as a learning exercise to practice:
- Go fundamentals (structs, methods, interfaces)
- File operations (JSON persistence)
- Test-Driven Development (TDD)
- Command-line argument parsing
- Error handling patterns

## Features

- ✅ Add tasks with automatic ID generation
- ✅ List all tasks or filter by status
- ✅ Update task descriptions
- ✅ Mark tasks as todo/in-progress/done
- ✅ Delete tasks
- ✅ Persistent storage using JSON
- ✅ Comprehensive test coverage

## Installation

### Prerequisites
- Go 1.21 or higher

### Build from source
```bash
# Clone the repository
git clone https://github.com/ayoblt/task-tracker.git
cd task-tracker

# Build the application
go build -o task-tracker

# (Optional) Install globally
go install
```

## Usage

### Add a new task
```bash
./task-tracker add "Buy groceries"
./task-tracker add "Learn Go"
```

### List tasks
```bash
# List all tasks
./task-tracker list

# Filter by status
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done
```

### Update a task
```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```

### Mark task status
```bash
# Mark as in progress
./task-tracker mark-in-progress 1

# Mark as done
./task-tracker mark-done 1

# Mark as todo
./task-tracker mark-todo 1
```

### Delete a task
```bash
./task-tracker delete 1
```

## Project Structure
```
task-tracker/
├── main.go              # Entry point and CLI routing
├── models/
│   ├── task.go          # Task struct and business logic
│   └── task_test.go     # Task tests
├── storage/
│   ├── storage.go       # JSON storage implementation
│   └── storage_test.go  # Storage tests
└── tasks.json           # Data file (created automatically)
```

## Architecture

### Storage Layer
The `storage` package provides a simple abstraction over JSON file operations:
```go
type Storage struct {
    filename string
}

// Main operations
func (s *Storage) Add(task models.Task) error
func (s *Storage) Load() ([]models.Task, error)
func (s *Storage) Update(task models.Task) error
func (s *Storage) Delete(id int) error
```

### Task Model
Tasks have the following structure:
```go
type Task struct {
    ID          int
    Description string
    Status      TaskStatus  // "todo", "in-progress", or "done"
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## Testing

This project follows Test-Driven Development (TDD) principles.

### Run all tests
```bash
go test ./...
```

### Run tests with coverage
```bash
go test -cover ./...
```

### Run tests verbosely
```bash
go test -v ./...
```

### Test Structure

- **Unit tests**: Test individual functions in isolation
- **Integration tests**: Test storage operations with actual files
- **Helper functions**: Reusable test utilities (`taskEqual`, `tasksEqual`)

Example test:
```go
func TestStorageAdd(t *testing.T) {
	testAddFile := "test_add.json"

	t.Cleanup(func() {
		os.Remove(testAddFile)
	})

	store := New(testAddFile)

	err := store.Add("Test Storage Add")

	want := models.Task{
		ID:          1,
		Description: "Test Storage Add",
		Status:      models.TaskStatusTodo,
	}

	if err != nil {
		t.Error("Error while adding to storage:", err)
	}

	savedTasks, err := store.Load()
	if err != nil {
		t.Error("Error loading test task file:", err)
	}

	if len(savedTasks) != 1 {
		t.Errorf("Error with saved tasks length: got %d, want 1", len(savedTasks))
	}
	assertTaskEqual(t, savedTasks[0], want)
}
```

## Development

### Adding a new command

1. Add handler in `cmd/handlers.go`:
```go
func handleMyCommand(storage *storage.Storage, args []string) {
    // Implementation
}
```

2. Add route in `main.go`:
```go
case "mycommand":
    cmd.handleMyCommand(store, args)
```

3. Write tests first (TDD):
```go
func TestMyCommand(t *testing.T) {
    // Test implementation
}
```

## What I Learned

Through building this project, I learned:

1. **File I/O in Go**: Reading/writing JSON, handling file existence
2. **TDD workflow**: Write test → watch fail → implement → refactor
3. **Error handling**: Using `fmt.Errorf` with `%w` for error wrapping
4. **Go testing**: Table-driven tests, test helpers, cleanup functions
5. **CLI design**: Argument parsing, command routing, user-friendly output
6. **Go idioms**: Struct methods, pointer receivers, interface design

## Future Improvements

- [ ] Add due dates for tasks
- [ ] Support for tags/categories
- [ ] Export to different formats (CSV, Markdown)
- [ ] Task priority levels
- [ ] SQLite backend option
- [ ] Color-coded terminal output
- [ ] Interactive mode
- [ ] Implement TUI(Terminal User Interface)

## License

MIT

## Acknowledgments

Built as a learning project following roadmap.sh's [Task Tracker project](https://roadmap.sh/projects/task-tracker).
