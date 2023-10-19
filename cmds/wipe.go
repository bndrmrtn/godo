package cmds

import (
	"bufio"
	"fmt"
	"github.com/bndrmrtn/godo/colors"
	"github.com/bndrmrtn/godo/models"
	"github.com/fatih/color"
	"os"
	"strings"
)

func sure() bool {
	c := color.New(color.FgRed)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(c.Sprint("Are you sure you want to wipe all your todos? [Yes/No]: "))
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(colors.Error("Something went wrong"))
		return false
	}
	text = strings.TrimSpace(strings.ToUpper(text))

	if text == "Y" || text == "YES" {
		return true
	}

	fmt.Println("Exiting...")
	return false
}

func Wipe(root *models.TodosRoot) {
	if sure() {
		root.Todos = []models.Todo{}
		fmt.Print(colors.Success("Todos successfully wiped"))
	}
}
