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
