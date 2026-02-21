package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tomhockett/task-cli/cli"
	"github.com/tomhockett/task-cli/task"
)

func main() {
	// Get home directory and create ~/.task-cli/tasks.db path
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get home directory:", err)
		os.Exit(1)
	}

	dbDir := filepath.Join(home, ".task-cli")
	// 0755 is Unix file permissions (rwxr-xr-x)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "failed to create data directory:", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(dbDir, "tasks.db")
	store, err := task.NewSQLiteStore(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open database:", err)
		os.Exit(1)
	}

	c := cli.NewCLI(store, os.Stdout)

	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
