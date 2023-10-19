package colors

import (
	"github.com/fatih/color"
	"strings"
)

func Success(message ...string) string {
	c := color.New(color.FgGreen)
	return c.Sprintln("✅ ", strings.Join(message, " "))
}

func Warning(message ...string) string {
	c := color.New(color.FgYellow)
	return c.Sprintln("⚠️ ", strings.Join(message, " "))
}

func Error(message ...string) string {
	c := color.New(color.FgRed)
	return c.Sprintln("❌ ", strings.Join(message, " "))
}

func Info(message ...string) string {
	c := color.New(color.FgBlue)
	return c.Sprintln("ℹ️ ", strings.Join(message, " "))
}
