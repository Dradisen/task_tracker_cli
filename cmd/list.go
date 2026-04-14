package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var list []string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LIST", cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
