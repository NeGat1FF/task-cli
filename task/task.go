package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	Todo       = 1
	InProgress = 2
	Done       = 3
)

type Task struct {
	ID          int
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var Tasks []Task

func AddTask(desc string) (int, error) {
	err := readFromFile()
	if err != nil {
		return 0, err
	}

	tsk := Task{1, desc, 1, time.Now(), time.Now()}

	if len(Tasks) == 0 {
		Tasks = append(Tasks, tsk)
	} else {
		tsk.ID = Tasks[len(Tasks)-1].ID + 1
		Tasks = append(Tasks, tsk)
	}

	return tsk.ID, writeToFile()
}

func GetTasks(n int) ([]Task, error) {
	tasks := make([]Task, 0, len(Tasks))

	err := readFromFile()
	if err != nil {
		return tasks, err
	}

	if n == 0 {
		return Tasks, nil
	}

	for _, task := range Tasks {
		if task.Status == n {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func UpdateTask(n int, desc string) error {
	err := readFromFile()
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("ID can't be less then 1")
	}
	found := false
	for i, task := range Tasks {
		if task.ID == n {
			found = true
			Tasks[i].Description = desc
			Tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID: %d not found", n)
	}

	return writeToFile()
}

func MarkTask(n, t int) error {
	err := readFromFile()
	if err != nil {
		return nil
	}

	if n < 1 {
		return errors.New("ID can't be less then 1")
	}
	found := false
	for i, task := range Tasks {
		if task.ID == n {
			found = true
			Tasks[i].Status = t
			Tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID: %d not found", n)
	}

	return writeToFile()
}

func DeleteTask(n int) error {
	err := readFromFile()
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("ID can't be less then 1")
	}
	found := false
	for i, task := range Tasks {
		if task.ID == n {
			found = true
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID: %d not found", n)
	}

	return writeToFile()
}

func getTimeAsString(t time.Time) string {
	return t.Format(time.DateTime)
}

func getStatusAsString(n int) string {
	switch n {
	case Todo:
		return "TODO"
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	}

	return ""
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d\nDescription: %s\nStatus: %s\nCreate At: %s\nUpdated At: %s\n", t.ID, t.Description, getStatusAsString(t.Status), getTimeAsString(t.CreatedAt), getTimeAsString(t.UpdatedAt))
}

func writeToFile() error {
	if len(Tasks) == 0 {
		return nil
	}
	data, err := json.Marshal(Tasks)
	if err != nil {
		return err
	}
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return file.Close()
}

func readFromFile() error {
	file, err := os.Open("tasks.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	data := make([]byte, fi.Size())
	n, err := file.Read(data)
	if err != nil {
		return err
	}

	if n != 0 {
		return json.Unmarshal(data, &Tasks)
	}

	return nil
}
