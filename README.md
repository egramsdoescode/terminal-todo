# Todo - Command-Line To-Do List

`todo` is a simple, cross-platform command-line to-do list application built with Go and Cobra. It helps you manage your tasks efficiently from the terminal.

## Features

- Add tasks to your to-do list
- List all your tasks
- Remove tasks from your list
- Cross-platform support (Linux, macOS, Windows)

## Installation

### Install via Go

You can install `todo` directly from the source using Go. Make sure you have Go installed and your `GOPATH/bin` is in your `PATH`.

```bash
go install github.com/egramsdoescode/todo@latest
```

## Commands
For simplicity, `todo` supports three easy-to-remember commands:

```bash
todo add [task] [time] # Add a task to your list
```
```bash
todo remove [task]     # Remove a task from your list
```
```bash
todo list             # View your list 
```

## Example Usage
```bash
# Add tasks
todo add "Write unit tests"
todo add "Fix bug in user authentication"
todo add "Prepare presentation slides"

# List tasks
todo list
# 1. Write unit tests
# 2. Fix bug in user authentication
# 3. Prepare presentation slides

# Mark the first task as done
todo done 1

# Remove the second task
todo rm 2

# List tasks again
todo list
# 1. Prepare presentation slides
```
