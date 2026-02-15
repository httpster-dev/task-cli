package task_test

import (
	"path/filepath"
	"testing"
)

// newTestStore creates a SQLiteStore backed by a temporary file.
// t.TempDir() is automatically cleaned up after the test — like DatabaseCleaner in Rails.
func newTestStore(t *testing.T) *SQLiteStore {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	store, err := NewSQLiteStore(path)
	if err != nil {
		t.Fatalf("failed to create SQLiteStore: %v", err)
	}
	return store
}

func TestSQLiteStore_Open(t *testing.T) {
	store := newTestStore(t)
	if store == nil {
		t.Fatal("expected a non-nil store")
	}
}

func TestSQLiteStore_Add(t *testing.T) {
	store := newTestStore(t)

	task, err := store.Add("Buy groceries")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.ID != 1 {
		t.Errorf("got ID %d, want 1", task.ID)
	}
	if task.Title != "Buy groceries" {
		t.Errorf("got title %q, want %q", task.Title, "Buy groceries")
	}
	if task.Status != StatusTodo {
		t.Errorf("got status %v, want StatusTodo", task.Status)
	}
}

func TestSQLiteStore_List(t *testing.T) {
	store := newTestStore(t)

	store.Add("Buy groceries")
	store.Add("Walk the dog")

	tasks, err := store.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("got %d tasks, want 2", len(tasks))
	}
	if tasks[0].Title != "Buy groceries" {
		t.Errorf("got title %q, want %q", tasks[0].Title, "Buy groceries")
	}
	if tasks[1].Title != "Walk the dog" {
		t.Errorf("got title %q, want %q", tasks[1].Title, "Walk the dog")
	}
}
