package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to list")
		return
	}
	fmt.Println(" Completed |  ID | Name")
	fmt.Println("---------------------------------------")
	for _, task := range tasks {
		status := " "

		if task.Completed {
			status = "✓ Yes"
		} else {
			status = "✖ No "
		}

		fmt.Printf("   %s   |   %d | %s\n", status, task.Id, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		Id:        GetNextId(tasks),
		Name:      name,
		Completed: false,
	}

	return append(tasks, newTask)
}

func GetNextId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].Id + 1
}

func ChageStatus(tasks []Task, id int, status bool) []Task {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Completed = status
			break
		}
	}
	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.Id == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
