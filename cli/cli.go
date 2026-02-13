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

func NewCLI(s task.TaskStore, b io.Writer) *CLI {
	return &CLI{
		store: s,
		out:   b,
	}
}

func (c *CLI) Run(s []string) error {
	if len(s) < 2 {
		return fmt.Errorf("missing task title")
	}
	taskCommand := string(s[0])
	taskTitle := string(s[1])

	switch taskCommand {
	case "add":
		_, err := c.store.Add(taskTitle)
		if err != nil {
			return err
		}
		fmt.Fprint(c.out, "Added task 1\n")
	}
	return nil
}
