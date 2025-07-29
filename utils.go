package main

import (
	"fmt"
	"sort"
)

// GetNextId возвращает первый свободный ID,
// не занятый не одной из существующих задач
func GetNextId() int {
	used := make(map[int]bool)

	for _, t := range tasks {
		used[t.ID] = true
	}

	for i := 1; ; i++ {
		if !used[i] {
			return i
		}
	}
}

// Sort сортирует задачи по возрастанию ID и сохраняет результат
func Sort() {
	LoadTasks()

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
	fmt.Println("Успешно отсортированно")

	SaveTasks()
}
