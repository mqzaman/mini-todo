package main

import "fmt"

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	taskid := len(td.tasks) + 1
	task := NewTask(taskid, taskName)
	td.tasks = append(td.tasks, task)
	fmt.Println("Added task successfully")
}
func (td *TodoList) ViewTask() {
	fmt.Println("============= TASKS ===============")
	for _, task := range td.tasks {
		strmark := ""
		if task.Done {
			strmark = "x"
		}
		fmt.Printf("[%s] Task #%d: %s\n", strmark, task.ID, task.Name)
	}
	fmt.Println("===================================")
}
func (td *TodoList) MarkAsDone(taskId int) {
	if taskId < 0 || taskId > len(td.tasks) {
		fmt.Println("Invalid task id")
		return
	}
	td.tasks[taskId-1].MarkAsDone()
	fmt.Println("Task marked as done")
}
