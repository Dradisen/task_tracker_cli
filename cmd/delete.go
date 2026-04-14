package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use:   "delete",
	Short: "Удаление объекта",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DELETE", cmd, args, id)
	},
}

func init() {
	rootCmd.AddCommand(&deleteCmd)
}
