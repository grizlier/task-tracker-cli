package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите команду: add/list/done/delete/sort")
		return
	}
	switch arg := os.Args[1]; arg {
	case "add", "Add":
		if len(os.Args) < 3 {
			fmt.Println("Укажите текст задачи после команды add")
			return
		}
		Add(os.Args[2])

	case "done", "Done":
		if len(os.Args) < 3 {
			fmt.Println("Укажите ID задачи после команды done")
			return
		}
		Done(os.Args[2])

	case "delete", "Delete":
		if len(os.Args) < 3 {
			fmt.Println("Укажите ID задачи после команды delete")
			return
		}
		Delete(os.Args[2])

	case "list", "List":
		List()

	case "sort", "Sort":
		Sort()

	default:
		fmt.Println("Неизвестная команда:", arg)
	}
}
