package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a to-do item
type Task struct {
	ID        int
	Content   string
	Completed bool
}

// TodoList represents a collection of tasks
type TodoList struct {
	tasks []Task
}

// AddTask adds a new task to the todo list
func (tl *TodoList) AddTask(content string) {
	newTask := Task{
		ID:        len(tl.tasks) + 1,
		Content:   content,
		Completed: false,
	}
	tl.tasks = append(tl.tasks, newTask)
}

// ListTasks prints all tasks in the todo list
func (tl *TodoList) ListTasks() {
	fmt.Println("Tasks:")
	for _, task := range tl.tasks {
		status := "x Not Completed 😥"
		if task.Completed {
			status = "✓ Completed 😁"
		}
		fmt.Printf("[%s] %s\n", status, task.Content)
	}
}

// CompleteTask marks a task as completed
func (tl *TodoList) CompleteTask(id int) {
	for i, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[i].Completed = true
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	todo := TodoList{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Complete task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		optionStr, _ := reader.ReadString('\n')
		fmt.Println(optionStr, "BEFORE selvin")
		optionStr = strings.TrimSpace(optionStr)
		fmt.Println(optionStr, "AFTER selvin")

		switch optionStr {
		case "1":
			fmt.Print("Enter task: ")
			taskContent, _ := reader.ReadString('\n')
			taskContent = strings.TrimSpace(taskContent)
			todo.AddTask(taskContent)
		case "2":
			todo.ListTasks()
		case "3":
			fmt.Print("Enter task ID to complete: ")
			var taskID int
			fmt.Scanln(&taskID)
			todo.CompleteTask(taskID)
		case "4":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
