package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/onglichang/todo-cli/todo"
)

const dataFile = "todos.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  todo add <task>         # Add a new task")
		fmt.Println("  todo list               # List pending tasks")
		fmt.Println("  todo list completed     # List completed tasks")
		fmt.Println("  todo done <taskID>      # Mark a task as completed")
		return
	}

	// Create list and load existing data
	list := todo.NewList()
	list.Load(dataFile)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name.")
			return
		}
		title := os.Args[2]
		list.Add(title)
		list.Save(dataFile)
		fmt.Println("Added:", title)

	case "list":
		var items []todo.Todo
		if len(os.Args) >= 3 && os.Args[2] == "completed" {
			items = list.Completed()
		} else {
			items = list.All()
		}

		if len(items) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, item := range items {
			status := " "
			if item.Done {
				status = "x"
			}
			fmt.Printf(
				"%d. [%s] %s (created: %s)\n",
				item.ID,
				status,
				item.Title,
				item.DateCreated.Format(time.RFC1123),
			)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if list.Complete(id) {
			list.Save(dataFile)
			fmt.Println("Task completed!")
		} else {
			fmt.Println("Task not found.")
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
