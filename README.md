# Todo CLI

A simple command-line **To-Do List** application written in **Go**.  
Tracks tasks with creation dates, marks tasks as done, and persists data in a text file (`todos.txt`).

---

## Features

- Add tasks with `todo add <task>`
- List **pending tasks** with `todo list`
- List **completed tasks** with `todo list completed`
- Mark tasks as done with `todo done <taskID>`
- Automatically saves tasks to `todos.txt`
- Completed tasks are moved to the bottom of the file

---

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/onglichang/todo-cli.git
cd todo-cli
