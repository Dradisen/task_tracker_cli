package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var text string

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ADD: ", text, args)
	},
}

func init() {
	createCmd.Flags().StringVarP(&text, "name", "n", "", "Добавление объекта/строки в список")

	rootCmd.AddCommand(createCmd)
}
