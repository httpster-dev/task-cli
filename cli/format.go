package cli

import (
	"fmt"

	"github.com/tomhockett/task-cli/task"
)

func FormatTaskTable(tasks []task.Task) string {
	if len(tasks) == 0 {
		return "No tasks\n"
	}
	var result string

	for _, t := range tasks {
		result += fmt.Sprintf("%d, %s, %s\n",
			t.ID,
			t.Title,
			t.Status)
	}

	return result
}
