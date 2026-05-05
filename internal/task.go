package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// id: A unique identifier for the task
// description: A short description of the task
// status: The status of the task (todo, in-progress, done)
// createdAt: The date and time when the task was created
// updatedAt: The date and time when the task was last updated
type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
type Status string
type ListTask []Task

const (
	New      Status = "new"
	Progress Status = "in-progress"
	Done     Status = "done"
)

// Валидация задачи
func (t *Task) Validate() error {
	if t.Id < 0 {
		return errors.New("Id must be more then zero")
	}

	allows_status := map[Status]bool{
		New:      true,
		Progress: true,
		Done:     true,
	}
	if !allows_status[t.Status] {
		return errors.New("Uncorrect Status")
	}

	return nil
}

// Загрузка списка
func (self *ListTask) LoadTasks() ([]byte, error) {
	bytes, err := os.ReadFile("list.json")
	if err != nil {
		return []byte{}, err
	}
	if err = json.Unmarshal(bytes, self); err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

// Добавление задачи
func (self *ListTask) Add(name string) (int, error) {

	next_seq := self.NextSeq()
	var task = Task{
		Id:          next_seq,
		Description: name,
		Status:      New,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	*self = append(*self, task)
	return next_seq, nil
}

// Фильтр списка по статусу
func (self *ListTask) Filter(status Status) ListTask {
	var filtered ListTask

	for _, task := range *self {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// Принтуха
func (self *ListTask) Print() {
	for _, task := range *self {
		fmt.Println(task.Id, task.Description, task.Status)
	}
}

// Обновление задачи
func (self *ListTask) Update(task_id int, description string) error {
	for i := range *self {
		task := &(*self)[i]
		if task.Id == task_id {
			task.Description = description
			task.UpdatedAt = time.Now().Unix()
		}
	}
	return nil
}

// Обновление статуса
func (self *ListTask) UpdateStatus(task_id int, status Status) {
	for i := range *self {
		task := &(*self)[i]
		if task.Id == task_id {
			task.Status = status
		}
	}
}

func (self *ListTask) Delete(task_id int) {
	// Удаление задачи
	sliced_tasks := *self
	first := true
	last := false

	if len(sliced_tasks) == 0 {
		return
	}

	for i := range *self {
		task := &(*self)[i]
		if i == len(sliced_tasks)-1 {
			last = true
		}
		if task.Id == task_id {
			if first && last {
				sliced_tasks = ListTask{}
			} else if first {
				sliced_tasks = sliced_tasks[i+1:]
			} else if last {
				sliced_tasks = sliced_tasks[:i]
			} else {
				sliced_tasks = append(sliced_tasks[:i], sliced_tasks[i+1:]...)
			}

		}
		first = false
	}
	*self = sliced_tasks
}

// Возвращает следующий номер последовательности
func (self *ListTask) NextSeq() int {
	var next_inc int = 0
	for _, task := range *self {
		if task.Id > next_inc {
			next_inc = task.Id
		}
	}

	return next_inc + 1
}

func (self *ListTask) Commit() error {
	// Фиксация данных в файл
	bytes, err := json.Marshal(*self)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("list.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
