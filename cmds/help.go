package cmds

import "fmt"

func Help(args []string) {
	commands := map[string]string{
		"add [...title]": "Create a new todo with the given title or titles separated by spaces.",
		"remove [...id]": "Remove a todo with the given ID or IDs separated by spaces.",
		"finish [...id]": "Set a todo as done with the given ID or IDs separated by spaces.",
		"show":           "Show all todos.",
		"clean":          "Removes all the done tasks.",
		"wipe":           "Removes all the tasks.",
		"helps":          "Shows this help menu.",
	}

	fmt.Println("GoDo - Help")
	fmt.Println("\nUsage:")
	fmt.Println("\tgodo <command> [arguments]")
	fmt.Println("\nCommands:")
	for cmd, desc := range commands {
		fmt.Println("\tgodo " + cmd + "\t" + desc)
	}
}
