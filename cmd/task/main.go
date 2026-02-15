package main

import (
	"fmt"
	"os"

	"github.com/tomhockett/task-cli/cli"
	"github.com/tomhockett/task-cli/task"
)

func main() {
	store := task.NewInMemoryTaskStore()
	c := cli.NewCLI(store, os.Stdout)

	// os.Args[1:] slices off the program name
	// os.Args[0] is always the binary name, so
	// os.Args[1:] gives you just the user's arguments
	if err := c.Run(os.Args[1:]); err != nil {
		// Errors go to os.Stderr and exit with code 1
		// the Unix convention for CLI tools
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
