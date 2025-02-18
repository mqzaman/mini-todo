package main

import "fmt"

type Task struct {
	ID   int
	Name string
	Done bool
}

func (t *Task) MarkAsDone() {
	if t.Done {
		fmt.Println("Task already done!")
	}
	t.Done = true
}
func NewTask(taskId int, taskName string) Task {
	return Task{ID: taskId, Name: taskName}
}
