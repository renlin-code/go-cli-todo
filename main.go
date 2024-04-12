package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/renlin-code/go-cli-todo/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}
	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Name your task:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTask(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide a task id")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id")
		}

		tasks = task.ChageStatus(tasks, id, true)
		task.SaveTask(file, tasks)
	case "uncomplete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide a task id")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id")
		}

		tasks = task.ChageStatus(tasks, id, false)
		task.SaveTask(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("You must provide a task id")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id")
		}

		tasks = task.DeleteTask(tasks, id)
		task.SaveTask(file, tasks)
	case "clear":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Are you sure you want to delete all your tasks? [y/n]")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == "y" {
			task.SaveTask(file, []task.Task{})
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`
	Usage: go-cli-todo [list|add|complete|uncomplete|delete|clear]
	
	list --------------- Show a list with all your tasks
	
	add ---------------- Open a prompt where you can add a name for a new task
	
	complete [id] ------ Change the status of a task to 'completed'

	uncomplete [id] ---- Change the status of a task to 'uncompleted'

	delete [id] -------- Delete a task

	clear -------------- Delete all your tasks
	`)
}
