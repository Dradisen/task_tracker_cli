package main

import (
	"encoding/json"
	"fmt"
	"os"
	"task_tracker_cli/cmd"
)

type Row struct {
	id          int    `json:id`
	description string `json:description`
	status      string `json:status`
	createdAt   int    `json:createdAt`
	updatedAt   int    `json:updatedAt`
}

var example = &Row{
	id:          1,
	description: "sadasd",
	status:      "test",
	createdAt:   12321,
	updatedAt:   12322,
}

func main() {

	filename := "list.json"

	bytes, err := os.ReadFile(filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	defer f.Close()

	results, err := json.Marshal(example)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	fmt.Println("JSON RESULT", string(results))

	n, err := f.Read(bytes)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	fmt.Println(n)

	// b := []byte{`dasd`}
	// f.Write(b)
	print("BYTES: ", example, string(bytes))

	cmd.Execute()
}
