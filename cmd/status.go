package cmd

import (
	"os"
	"strconv"
	"task_tracker_cli/internal"

	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Помечает задачу в работу",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var list_task internal.ListTask
		list_task.LoadTasks()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		list_task.UpdateStatus(id, internal.Progress)
		list_task.Commit()
	},
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Помечает задачу выполненной",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var list_task internal.ListTask
		list_task.LoadTasks()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			os.Exit(1)
		}
		list_task.UpdateStatus(id, internal.Done)
		list_task.Commit()
	},
}
