package util

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintGoreLogo() {
	fmt.Println()
	color.Red(" .d8888b.        .d88888b.       8888888b.       8888888888 ")
	color.Red("d88P  Y88b      d88P   Y88b      888   Y88b      888        ")
	color.Red("888    888      888     888      888    888      888        ")
	color.Red("888             888     888      888   d88P      8888888    ")
	color.Red("888  88888      888     888      8888888P        888        ")
	color.Red("888    888      888     888      888 T88b        888        ")
	color.Red("Y88b  d88P      Y88b. .d88P      888  T88b       888        ")
	color.Red("  Y8888P88        Y88888P        888   T88b      8888888888 ")
}
