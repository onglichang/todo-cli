# Todo CLI

A simple command-line To-Do List application written in Go.  
Tracks tasks with creation dates, marks tasks as done, and persists data in a text file (`todos.txt`).

## Features

- Add tasks with `todo add <task>`
- List pending tasks with `todo list`
- List completed tasks with `todo list completed`
- Mark tasks as done with `todo done <taskID>`
- Automatically saves tasks to `todos.txt`
- Completed tasks are moved to the bottom of the file

## Prerequisites

### Go Installation

Check if Go is installed:

```
go version
```

If not installed, download Go from:

https://go.dev/dl/

After installation, verify GOPATH:

```
echo $GOPATH
```

Ensure `~/go/bin` is in your PATH:

```
echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshrc
source ~/.zshrc
```

## Installation

### 1. Clone the repository

```
git clone https://github.com/onglichang/todo-cli.git
cd todo-cli
```

### 2. Install the CLI

```
go install
```

After this, ensure `~/go/bin` is in your PATH (see prerequisites section).

You should then be able to run:

```
todo list
```

## Usage

### Add a new task

```
todo add "Buy groceries"
```

### List all pending tasks

```
todo list
```

### List completed tasks

```
todo list completed
```

### Mark a task as done

```
todo done <taskID>
```

## Data Storage

All tasks are stored in a file named `todos.txt`.

Format:

```
[ ] Buy groceries | 2025-11-20T19:30:00Z
[x] Walk the dog  | 2025-11-19T11:00:00Z
```

- `[ ]` indicates a pending task  
- `[x]` indicates a completed task  
- Completed tasks are automatically moved to the bottom of the file

## Project Structure

```
todo-cli/
├── main.go
├── todo/
│   ├── todo.go
├── todos.txt
└── README.md
```

## Contributing

Pull requests are welcome.
