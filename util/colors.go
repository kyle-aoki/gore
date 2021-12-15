package util

import "github.com/fatih/color"

var Green func(a ...interface{}) string
var Blue func(a ...interface{}) string

func SetColors() {
	Green = color.New(color.FgGreen).SprintFunc()
	Blue = color.New(color.FgCyan).SprintFunc()
}
