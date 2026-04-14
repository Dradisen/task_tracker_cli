package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var id int
var object string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновление объекта/строки",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UPDATE", cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
