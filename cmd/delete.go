package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use:   "delete",
	Short: "Удаление объекта",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var list_task ListTask
		list_task.LoadTasks()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		list_task.Delete(id)
		list_task.Commit()
	},
}

func init() {
	rootCmd.AddCommand(&deleteCmd)
}
