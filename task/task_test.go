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

	assertEqual(t, tk.ID, 1)
	assertEqual(t, tk.Title, "Buy groceries")
	assertEqual(t, tk.Status, task.StatusTodo)
	assertEqual(t, tk.Priority, task.PriorityMedium)
	assertEqual(t, tk.CreatedAt, now)
	assertEqual(t, tk.CompletedAt, nil)
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
		assertEqual(t, got, tt.want)
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
		assertEqual(t, got, tt.want)
	}
}
