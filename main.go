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
		fmt.Println("Usage: todo [add|list|done] <task>")
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
		for _, item := range list.All() {
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
