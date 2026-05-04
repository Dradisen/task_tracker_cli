package main

import (
	"encoding/json"
	"fmt"
	"os"
	"task_tracker_cli/cmd"
)

// type Row struct {
// 	id          int    `json:id`
// 	description string `json:description`
// 	status      string `json:status`
// 	createdAt   int    `json:createdAt`
// 	updatedAt   int    `json:updatedAt`
// }

// Запись в файл задачи
func write_task(file *os.File, task *cmd.Task) (int, error) {

	results, err := json.Marshal(task)
	if err != nil {
		return -1, err
	}

	len, err := file.WriteString(string(results) + "\n")

	if err != nil {
		return -1, err
	}
	return len, nil
}

func main() {
	var lst cmd.ListTask

	lst.LoadTasks()
	lst.Add(cmd.Task{Id: 1})
	lst.Add(cmd.Task{Id: 2})
	lst.Add(cmd.Task{Id: 3})

	err := lst.WriteToFile()

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Println("???", lst)

	cmd.Execute()
}
