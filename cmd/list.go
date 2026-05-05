package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список",
	Run: func(cmd *cobra.Command, args []string) {
		var task_list ListTask
		task_list.LoadTasks()
		if len(args) > 0 {
			switch Status(args[0]) {

			case New, Progress, Done:
				task_list = task_list.Filter(Status(args[0]))
			default:
				return
			}
		}
		task_list.Print()
	},
}
