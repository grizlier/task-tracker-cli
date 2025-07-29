package main

import (
	"fmt"
	"strconv"
	"time"
)

// Task представляет собой задачу в списке дел
type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

var tasks []Task

// Add добавляет новую задачу с текстом arg в список и сохраняет его
func Add(arg string) {
	LoadTasks()

	id := GetNextId()
	task := Task{
		ID:        id,
		Text:      arg,
		Status:    false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, task)
	fmt.Printf("Задача добавлена: [%d] %s\n", id, arg)

	SaveTasks()
}

// List выводит список всех задач в консоль
func List() {
	LoadTasks()

	fmt.Println("[Id]\t Статус\t Описание")
	for _, t := range tasks {
		fmt.Printf("[%d]\t %t\t %s \n", t.ID, t.Status, t.Text)
	}
}

// Done отмечает задачу с выбранным ID как выполненную
func Done(arg string) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Введен неправильный ID")
		return
	}
	LoadTasks()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = true
		}
	}
	fmt.Printf("Запись %d была отмечена как выполненная \n", id)

	SaveTasks()
}

// Delete удаляет задачу с выбранным ID из списка
func Delete(arg string) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Введен неправильный ID")
		return
	}
	LoadTasks()

	newTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			fmt.Printf("Запись [%d] удалена \n", id)
		}
	}
	tasks = newTasks

	SaveTasks()
}
