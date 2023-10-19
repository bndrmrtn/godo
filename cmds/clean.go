package cmds

import (
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
)

func Clean(root *models.TodosRoot) {
	cleanCount := 0
	for id, todo := range root.Todos {
		if todo.IsDone {
			cleanCount++
			root.Todos = append(root.Todos[:id], root.Todos[id+1:]...)
		}
	}
	if cleanCount == 0 {
		fmt.Print(colors.Info("Nothing to clean"))
	} else {
		s := ""
		if cleanCount > 1 {
			s = "s"
		}
		fmt.Print(colors.Success(fmt.Sprintf("%v todo%v successfully cleaned successfully", cleanCount, s)))
	}
	Show(root)
}
