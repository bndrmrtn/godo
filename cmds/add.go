package cmds

import (
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"time"
)

func Add(root *models.TodosRoot, args []string) {
	for _, arg := range args {
		todo := models.Todo{
			Title:      arg,
			IsDone:     false,
			CreatedAt:  time.Now().UTC(),
			FinishedAt: time.Now().UTC(),
		}
		root.Todos = append(root.Todos, todo)
		fmt.Print(colors.Success("Todo added successfully"))
	}
	Show(root)
}
