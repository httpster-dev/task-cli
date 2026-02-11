package task_test

import (
	"testing"
	"time"

	"github.com/tomhockett/task-cli/task"
)

func TestTaskCreation(t *testing.T) {
	now := time.Now()
	tk := task.Task{
		ID:        1,
		Title:     "Buy groceries",
		Status:    task.StatusTodo,
		Priority:  task.PriorityMedium,
		CreatedAt: now,
	}

	if tk.ID != 1 {
		t.Errorf("got ID %d, want 1", tk.ID)
	}
	if tk.Title != "Buy groceries" {
		t.Errorf("got Title %q, want %q", tk.Title, "Buy groceries")
	}
	if tk.Status != task.StatusTodo {
		t.Errorf("got Status %d, want StatusTodo", tk.Status)
	}
	if tk.Priority != task.PriorityMedium {
		t.Errorf("got Priority %d, want PriorityMedium", tk.Priority)
	}
	if tk.CreatedAt != now {
		t.Errorf("got CreatedAt %v, want %v", tk.CreatedAt, now)
	}
	if tk.CompletedAt != nil {
		t.Error("expected CompletedAt to be nil for a new task")
	}
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		status task.Status
		want   string
	}{
		{task.StatusTodo, "todo"},
		{task.StatusDone, "done"},
	}

	for _, tt := range tests {
		got := tt.status.String()
		if got != tt.want {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}

func TestPriorityString(t *testing.T) {
	tests := []struct {
		priority task.Priority
		want     string
	}{
		{task.PriorityLow, "low"},
		{task.PriorityMedium, "medium"},
		{task.PriorityHigh, "high"},
	}

	for _, tt := range tests {
		got := tt.priority.String()
		if got != tt.want {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}
