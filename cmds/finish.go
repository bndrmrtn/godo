package cmds

import (
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"strconv"
	"time"
)

func Finish(root *models.TodosRoot, args []string) {
	for _, strId := range args {
		id, err := strconv.Atoi(strId)
		if err != nil {
			fmt.Print(colors.Error("The parameter \"" + strId + "\" should be an identifier"))
		} else {
			id--
			for todoId, _ := range root.Todos {
				if id == todoId {
					if root.Todos[todoId].IsDone == true {
						fmt.Print(colors.Warning("Todo", strId, "is already done"))
					} else {
						root.Todos[todoId].FinishedAt = time.Now()
						root.Todos[todoId].IsDone = true
						fmt.Print(colors.Success("Todo", strId, "set as done"))
						break
					}
				}
			}
		}
	}
	Show(root)
}
