package task_test

import (
	"errors"
	"testing"

	"github.com/tomhockett/task-cli/task"
)

func TestInMemoryStore_AddAndList(t *testing.T) {
	store := task.NewInMemoryTaskStore()

	// Add first task
	t1, err := store.Add("Buy groceries")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if t1.ID != 1 {
		t.Errorf("first task ID: got %d, want 1", t1.ID)
	}
	if t1.Title != "Buy groceries" {
		t.Errorf("got title %q, want %q", t1.Title, "Buy groceries")
	}
	if t1.Status != task.StatusTodo {
		t.Errorf("got status %v, want StatusTodo", t1.Status)
	}

	// Add second task — ID auto-increments
	t2, err := store.Add("Walk the dog")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if t2.ID != 2 {
		t.Errorf("second task ID: got %d, want 2", t2.ID)
	}

	// List returns both
	tasks, err := store.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("got %d tasks, want 2", len(tasks))
	}
}

func TestInMemoryStore_ListEmpty(t *testing.T) {
	store := task.NewInMemoryTaskStore()

	tasks, err := store.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("got %d tasks, want 0", len(tasks))
	}
}

func TestInMemoryStore_Complete(t *testing.T) {
	store := task.NewInMemoryTaskStore()
	store.Add("Buy groceries")

	err := store.Complete(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tasks, _ := store.List()
	if tasks[0].Status != task.StatusDone {
		t.Errorf("got status %v, want StatusDone", tasks[0].Status)
	}
	if tasks[0].CompletedAt == nil {
		t.Error("expected CompletedAt to be set after completing")
	}
}

func TestInMemoryStore_CompleteNotFound(t *testing.T) {
	store := task.NewInMemoryTaskStore()

	err := store.Complete(999)
	if !errors.Is(err, task.ErrTaskNotFound) {
		t.Errorf("got error %v, want ErrTaskNotFound", err)
	}
}
