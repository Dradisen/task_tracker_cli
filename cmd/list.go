package cmd

import (
	"task_tracker_cli/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список",
	Run: func(cmd *cobra.Command, args []string) {
		var task_list internal.ListTask
		task_list.LoadTasks()
		if len(args) > 0 {
			switch internal.Status(args[0]) {

			case internal.New, internal.Progress, internal.Done:
				task_list = task_list.Filter(internal.Status(args[0]))
			default:
				return
			}
		}
		task_list.Print()
	},
}
