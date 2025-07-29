package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadTasks загружает список задач из файла "tasks.json" в память
func LoadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Ошибка загрузки: ", err)
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&tasks)
}

// SaveTasks сохраняет текущий список задач в файл "tasks.json"
func SaveTasks() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Ошибка сохранения: ", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)
}
