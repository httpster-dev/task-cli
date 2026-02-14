package cli

import (
	"fmt"
	"io"

	"github.com/tomhockett/task-cli/task"
)

type CLI struct {
	store task.TaskStore
	out   io.Writer
}

func NewCLI(store task.TaskStore, out io.Writer) *CLI {
	return &CLI{
		store: store,
		out:   out,
	}
}

func (c *CLI) Run(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("usage: task <command> [args]")
	}
	taskCommand := s[0]

	switch taskCommand {
	case "add":
		if len(s) < 2 {
			return fmt.Errorf("missing task title")
		}
		taskTitle := s[1]
		t, err := c.store.Add(taskTitle)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.out, "Added task %d\n", t.ID)
	case "list":
		tasks, err := c.store.List()
		if err != nil {
			return err
		}
		output := FormatTaskTable(tasks)
		fmt.Fprint(c.out, output)
	default:
		return fmt.Errorf("unknown command %q", taskCommand)
	}
	return nil
}
