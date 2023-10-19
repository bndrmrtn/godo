package main

import (
	"encoding/json"
	"fmt"
	"github.com/bndrmrtn/godo/cmds"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"os"
)

const StorageFile = "./data/store.json"

func loadTodos() models.TodosRoot {
	var root models.TodosRoot
	data, err := os.ReadFile(StorageFile)
	if err != nil {
		fmt.Print(colors.Info("Previous save cannot be loaded:", err.Error()))
		return root
	}
	if err := json.Unmarshal(data, &root); err != nil {
		fmt.Print(colors.Info("Previous save cannot be loaded:", err.Error()))
	}
	return root
}

func saveTodos(root models.TodosRoot) {
	todos, err := json.Marshal(root)
	if err != nil {
		panic(colors.Error("Failed to backup the progress:", err.Error()))
	}
	if err := os.WriteFile(StorageFile, todos, 0644); err != nil {
		panic(colors.Error("Failed to backup the progress:", err.Error()))
	}
}

func welcome() {
	fmt.Println("Welcome to GoDo ;)")
	fmt.Println("To see all the functionality run \"godo help\"")
}

func main() {
	var method string
	var args []string
	if len(os.Args) > 1 {
		method = os.Args[1]
		if len(os.Args) > 2 {
			args = os.Args[2:]
		}
	}

	todoRoot := loadTodos()

	switch method {
	case "add":
		cmds.Add(&todoRoot, args)
	case "remove":
		cmds.Remove(&todoRoot, args)
	case "finish":
		cmds.Finish(&todoRoot, args)
	case "clean":
		cmds.Clean(&todoRoot)
	case "wipe":
		cmds.Wipe(&todoRoot)
	case "show":
		cmds.Show(&todoRoot)
	case "help":
		cmds.Help(args)
	default:
		if len(os.Args) == 1 {
			welcome()
		} else {
			fmt.Print(colors.Error("Invalid command", "\""+method+"\""))
			fmt.Print(colors.Info("Try run \"godo help\" for a list of commands"))
			os.Exit(1)
		}
	}

	saveTodos(todoRoot)
}
