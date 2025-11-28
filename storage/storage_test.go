package storage

import (
	"os"
	"testing"

	"github.com/ayoblt/task-tracker/models"
)

const testDataFile = "tasks_test.json"
const nonExistentFile = "file_not_exists.json"

func TestNewStorage(t *testing.T) {
	testNewFile := "test_new.json"

	t.Cleanup(func() { os.Remove(testNewFile) })

	store := New(testNewFile)

	err := store.Add("Test")

	if err != nil {
		t.Error("Error when testing returned storage", err)
	}
}

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

func TestStorageUpdate(t *testing.T) {
	testUpdateFile := "test_update.json"
	t.Cleanup(func() { os.Remove(testUpdateFile) })

	store := New(testUpdateFile)

	store.Add("Original")

	want := models.Task{ID: 1, Description: "Updated", Status: models.TaskStatusTodo}

	err := store.Update(1, "Updated")
	if err != nil {
		t.Fatal("Update failed:", err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Error("failed to load test_update.json:")
	}
	assertTaskEqual(t, loadedTasks[0], want)
}

func TestStorageDelete(t *testing.T) {
	testDeleteFile := "test_delete.json"
	t.Cleanup(func() { os.Remove(testDeleteFile) })

	store := New(testDeleteFile)

	store.Add("Should be deleted")

	err := store.Delete(1)
	if err != nil {
		t.Error(err)
	}

	loadedTasks, err := store.Load()
	want := []models.Task{}

	if err != nil {
		t.Errorf("failed to load %s", testDeleteFile)
	}
	if len(loadedTasks) != 0 {
		t.Errorf("Task not deleted got %q, want %q", loadedTasks, want)
	}
}

func TestStorageMarkInProgress(t *testing.T) {
	testMarkFile := "test_mark_in_progres.json"
	t.Cleanup(func() { os.Remove(testMarkFile) })

	store := New(testMarkFile)

	store.Add("Should be Marked In-progress")

	err := store.MarkInProgress(1)
	if err != nil {
		t.Error(err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Error("failed to load file")
	}

	if loadedTasks[0].Status != models.TaskStatusInProgress {
		t.Errorf("Status mismatch: got %s, want %s", loadedTasks[0].Status, models.TaskStatusInProgress)
	}
}

func TestStorageMarkDone(t *testing.T) {
	testMarkFile := "test_mark_done.json"
	t.Cleanup(func() { os.Remove(testMarkFile) })

	store := New(testMarkFile)

	store.Add("Should be Marked done")

	err := store.MarkDone(1)
	if err != nil {
		t.Error(err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Error("failed to load file")
	}

	if loadedTasks[0].Status != models.TaskStatusDone {
		t.Errorf("Status mismatch: got %s, want %s", loadedTasks[0].Status, models.TaskStatusInProgress)
	}
}

func TestStorageMarkTodo(t *testing.T) {
	testMarkFile := "test_mark_todo.json"
	t.Cleanup(func() { os.Remove(testMarkFile) })

	store := New(testMarkFile)

	store.Add("Should be Marked Todo")

	err := store.MarkTodo(1)
	if err != nil {
		t.Error(err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Error("failed to load file")
	}

	if loadedTasks[0].Status != models.TaskStatusTodo {
		t.Errorf("Status mismatch: got %s, want %s", loadedTasks[0].Status, models.TaskStatusInProgress)
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	testSaveAndLoadFile := "test_save_and_load.json"

	t.Cleanup(func() {
		os.Remove(testSaveAndLoadFile)
	})

	store := New(testSaveAndLoadFile)

	tasks := []models.Task{}
	newTask := models.CreateTask("Test for Adding", tasks)
	tasks = append(tasks, newTask)

	err := store.Save(tasks)
	if err != nil {
		t.Errorf("Error while saving tasks: %q", err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Errorf("Error while loading tasks: %q", err)
	}

	assertTasksEqual(t, loadedTasks, tasks)
}

func TestStorageUpdateNonExistent(t *testing.T) {
	testFile := "test_update_nonexistent.json"
	t.Cleanup(func() { os.Remove(testFile) })

	store := New(testFile)

	// Add two tasks
	store.Add("Task 1")
	store.Add("Task 2")

	// Try to update task that doesn't exist
	err := store.Update(999, "Ghost task")
	if err == nil {
		t.Error("Update happend on non existent")
	}
	// What do you EXPECT to happen?
	// Write your expectations here before running the test!

	// Load and check
	tasks, _ := store.Load()
	t.Logf("Number of tasks: %d", len(tasks))
	for _, task := range tasks {
		t.Logf("Task: ID=%d, Desc=%s", task.ID, task.Description)
	}
}

func TestLoadTasksFileNotExists(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(nonExistentFile)
	})

	store := New(nonExistentFile)

	got, err := store.Load()

	if err != nil {
		t.Error("Test Tasks File Failed:", err)
	}

	if len(got) > 0 {
		t.Error("Test Tasks File not Exists failed: File Exists")
	}

}

func assertTaskEqual(t testing.TB, got, want models.Task) {
	t.Helper()

	if got.ID != want.ID {
		t.Errorf("Task: ID mismatch: got %d, want %d", got.ID, want.ID)
	}

	if got.Description != want.Description {
		t.Errorf("Task: Description mismatch: got '%s', want '%s'", got.Description, want.Description)

	}

	if got.Status != want.Status {
		t.Errorf("Task: Status mismatch: got '%s', want '%s'", got.Status, want.Status)
	}
}

func assertTasksEqual(t testing.TB, got, want []models.Task) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("Length mismatch: got %d, want %d", len(got), len(want))
	}

	for i := range got {
		assertTaskEqual(t, got[i], want[i])
	}
}

// func TestTimeJSON(t *testing.T) {
// 	now := time.Now()
// 	fmt.Println("Before JSON:", now)

// 	data, _ := json.Marshal(now)
// 	fmt.Println("JSON:", string(data))

// 	var parsed time.Time
// 	json.Unmarshal(data, &parsed)
// 	fmt.Println("After JSON:", parsed)

// 	fmt.Println("Are they equal?", now == parsed)
// }
