# task-cli

`task-cli` is a command-line interface tool for managing tasks. It allows you to add, update, delete, and list tasks, as well as mark tasks with various statuses. This tool helps you keep track of your tasks efficiently through simple and intuitive commands.

## Features

- **Add tasks** with descriptions.
- **List all tasks** or filter by status (TODO, In Progress, Done).
- **Update task descriptions**.
- **Mark tasks** as TODO, In Progress, or Done.
- **Delete tasks** by their ID.

## Installation

To use `task-cli`, clone the repository and build the application:

```sh
git clone https://github.com/yourusername/task-cli.git
cd task-cli
go build -o task-cli
```
After building, you can use the `task-cli` binary from the command line.

# Usage
## Add a New Task
```sh
task-cli add [task description]
```
Example:
```sh
task-cli add "Complete project report"
```

## List Tasks
List all tasks:
```sh
task-cli list
```
List tasks with a specific status:
```sh
task-cli list todo
task-cli list done
task-cli list in-progress
```

## Update a Task
Update the description of a task by its ID:
```sh
task-cli update [task ID] "[new description]"
```
Example:
```sh
task-cli update 1 "Complete the final review"
```
## Mark a Task
Mark a task as done:
```sh
task-cli mark done [task ID]
```
Mark a task as in progress:
```sh
task-cli mark in-progress [task ID]
```
## Delete a Task
Delete a task by its ID:
```sh
task-cli delete [task ID]
```

## Example
```sh
# Add a new task
task-cli add "Buy groceries"

# List all tasks
task-cli list

# Update a task
task-cli update 2 "Buy groceries and cook dinner"

# Mark a task as done
task-cli mark done 2

# Delete a task
task-cli delete 2
```
