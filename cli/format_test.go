package cli_test

import (
	"testing"
	"time"

	"github.com/tomhockett/task-cli/cli"
	"github.com/tomhockett/task-cli/task"
)

func TestFormatTaskTable(t *testing.T) {
	now := time.Now()
	tasks := []task.Task{
		{ID: 1, Title: "Buy groceries", Status: task.StatusTodo, Priority: task.PriorityMedium, CreatedAt: now},
		{ID: 2, Title: "Walk the dog", Status: task.StatusDone, Priority: task.PriorityHigh, CreatedAt: now},
	}

	got := cli.FormatTaskTable(tasks)

	// Should contain each task's ID, title, and status
	for _, want := range []string{"1", "Buy groceries", "todo", "2", "Walk the dog", "done"} {
		if !contains(got, want) {
			t.Errorf("output missing %q:\n%s", want, got)
		}
	}
}

func TestFormatTaskTable_Empty(t *testing.T) {
	got := cli.FormatTaskTable([]task.Task{})

	want := "No tasks\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
