package cli

import (
	"fmt"
	"strings"

	"github.com/tomhockett/task-cli/task"
)

func FormatTaskTable(tasks []task.Task) string {
	if len(tasks) == 0 {
		return "No tasks\n"
	}
	var sb strings.Builder
	for _, t := range tasks {
		fmt.Fprintf(&sb, "%d, %s, %s\n", t.ID, t.Title, t.Status)
	}
	return sb.String()
}
