package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Программа-задачник",
	Long:  "Проект «Трекер задач» предназначен для отслеживания и управления вашими задачами",
}

func Execute() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
