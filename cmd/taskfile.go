package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go.ua-ecm.com/chaki/tasks"
)

func addTaskConfigFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("tasks-file", "t", "tasks.yaml", "file for loading task list")
}

func loadTaskConfig(cmd *cobra.Command) (*tasks.Config, error) {
	tasksFile, err := cmd.PersistentFlags().GetString("tasks-file")
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stderr, "using task file: %s\n", tasksFile)

	file, err := os.Open(tasksFile)
	if err != nil {
		return nil, err
	}

	return tasks.NewConfig(file)
}
