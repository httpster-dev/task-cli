package cli_test

import (
	"bytes"
	"testing"

	"github.com/tomhockett/task-cli/cli"
	"github.com/tomhockett/task-cli/task"
)

// newTestCLI wires up a CLI with an in-memory store and a buffer to capture output.
// Like injecting a StringIO in a Rails service test — no real I/O needed.
func newTestCLI() (*cli.CLI, *bytes.Buffer) {
	store := task.NewInMemoryTaskStore()
	buf := &bytes.Buffer{}
	return cli.NewCLI(store, buf), buf
}

func TestCLI_Add(t *testing.T) {
	c, buf := newTestCLI()

	err := c.Run([]string{"add", "Buy groceries"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Output should confirm the task was created
	got := buf.String()
	want := "Added task 1\n"
	if got != want {
		t.Errorf("got output %q, want %q", got, want)
	}
}

func TestCLI_Add_MissingTitle(t *testing.T) {
	c, _ := newTestCLI()

	err := c.Run([]string{"add"})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCLI_List(t *testing.T) {
	c, buf := newTestCLI()

	c.Run([]string{"add", "Buy groceries"})
	c.Run([]string{"add", "Walk the dog"})

	buf.Reset() // clear add output, only care about list output
	err := c.Run([]string{"list"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	if !contains(got, "Buy groceries") {
		t.Errorf("output missing %q:\n%s", "Buy groceries", got)
	}
	if !contains(got, "Walk the dog") {
		t.Errorf("output missing %q:\n%s", "Walk the dog", got)
	}
	if !contains(got, "1") {
		t.Errorf("output missing task ID 1:\n%s", got)
	}
}

func TestCLI_List_Empty(t *testing.T) {
	c, buf := newTestCLI()

	err := c.Run([]string{"list"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	want := "No tasks\n"
	if got != want {
		t.Errorf("got output %q, want %q", got, want)
	}
}

// contains is a small helper to avoid importing strings in the test file.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsAt(s, substr))
}

func containsAt(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
