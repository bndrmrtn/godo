package cmds

import (
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"strconv"
)

func Remove(root *models.TodosRoot, args []string) {
	for _, strId := range args {
		id, err := strconv.Atoi(strId)
		if err != nil {
			fmt.Print(colors.Error("The parameter \"" + strId + "\" should be an identifier"))
		} else {
			id--
			for todoId, _ := range root.Todos {
				if id == todoId {
					root.Todos = append(root.Todos[:id], root.Todos[id+1:]...)
					fmt.Print(colors.Success("Todo", strId, "removed successfully"))
					break
				}
			}
		}
	}
	Show(root)
}
