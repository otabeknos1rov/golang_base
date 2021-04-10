package data

import (
	"errors"
	"fmt"
)

type Task struct {
	Id     int
	Title  string
	Status bool
}

var taskList []Task

type Task1 interface {
	Add(t *Task) ([]Task, error)
	Get(id int) *Task
	Update(t Task) (Task, error)
	List() *TaskList
	Delete(id int) ([]Task, error)
}

type TaskList struct {
	Tasks []*Task
}

func NewTask() Task1 {
	return &TaskList{}
}

func (task *TaskList) Add(t *Task) ([]Task, error) {
	if t.Id == 0 {
		return nil, errors.New("Id field is empty ")
	}
	if t.Title == "" {
		return nil, errors.New("Title field is empty ")
	}

	task.Tasks = append(task.Tasks, t)

	return taskList, nil
}

func (task *TaskList) Get(id int) *Task {

	foundTask := new(Task)

	if id != 0 {
		fmt.Printf("Id is null")
	}

	for i := 0; i < len(task.Tasks); i++ {
		if task.Tasks[i].Id == id {
			foundTask = task.Tasks[i]
		}
	}

	return foundTask
}

func (task *TaskList) Update(newTask Task) (Task, error) {

	if newTask.Id != 0 {
		fmt.Printf("Id is null")
	}
	if newTask.Title != "" {
		fmt.Printf("Title is empty")
	}

	foundTask := new(Task)

	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == newTask.Id {
			foundTask = &taskList[i]
		}
	}

	if foundTask == nil {
		return *foundTask, fmt.Errorf("Task not found")
	}

	return *foundTask, nil
}

func (task *TaskList) List() *TaskList {
	return task
}

func (task *TaskList) Delete(id int) ([]Task, error) {

	if id != 0 {
		fmt.Printf("Id is null")
	}

	var list []Task

	for i := 0; i < len(taskList); i++ {

		if taskList[i].Id != id {
			list = append(list, taskList[i])
		}
	}

	if len(list) == 0 {
		return list, fmt.Errorf("Task not found")
	}

	taskList = list

	return taskList, nil
}
