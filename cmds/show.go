package cmds

import (
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"github.com/dustin/go-humanize"
)

func Show(root *models.TodosRoot) {
	if len(root.Todos) == 0 {
		fmt.Print(colors.Success("Currently have nothing to do"))
	} else {
		fmt.Println("Things to do:")
		for index, todo := range root.Todos {
			emoji := "⌚"
			if todo.IsDone {
				emoji = "✅"
			}
			fmt.Printf("%v  %v. %v\n", emoji, index+1, todo.Title)
			if todo.IsDone {
				fmt.Printf("\t• It took %v\n", humanize.RelTime(todo.FinishedAt, todo.CreatedAt, "", ""))
			}
		}
	}
}
