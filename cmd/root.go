package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// id: A unique identifier for the task
// description: A short description of the task
// status: The status of the task (todo, in-progress, done)
// createdAt: The date and time when the task was created
// updatedAt: The date and time when the task was last updated

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListTask []Task

func (self *ListTask) LoadTasks() ([]byte, error) {
	// Загрузка списка
	bytes, err := os.ReadFile("list.json")
	if err != nil {
		return []byte{}, err
	}
	err = json.Unmarshal(bytes, self)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (self *ListTask) Add(task Task) error {
	// Добавление задачи
	*self = append(*self, task)
	return nil
}

func (self *ListTask) Update(task_id int) error {
	// Обновление задачи
	return nil
}

func (self *ListTask) Delete(task_id int) error {
	// Удаление задачи
	return nil
}

func (self *ListTask) WriteToFile() error {
	// Фиксация данных в файл
	bytes, err := json.Marshal(*self)
	if err != nil {
		return err
	}
	err = os.WriteFile("list.json", bytes, os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Программа-задачник",
	Long:  "Длинное описание программы",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}
