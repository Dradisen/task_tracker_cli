package cmd

import (
	"fmt"
	"os"
	"task_tracker_cli/internal"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var list_task internal.ListTask
		list_task.LoadTasks()
		id, err := list_task.Add(args[0])
		if err != nil {
			os.Exit(1)
		}
		list_task.Commit()
		fmt.Printf("Task added successfully (ID: %d)\n", id)

	},
}
