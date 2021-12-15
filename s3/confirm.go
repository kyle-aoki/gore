package s3

import (
	"fmt"
	"gore/input"
	"gore/model"
	"os"

	"github.com/fatih/color"
)

func ConfirmUpload(newRelease model.Release) {
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println()
	fmt.Printf("Type 'confirm' to CONFIRM release: %s", green(newRelease.QualifiedName()))
	fmt.Println()
	
	textInput := input.WaitForInputText()

	if textInput != "confirm" {
		fmt.Println("User did not confirm, exiting program.")
		os.Exit(0)
	}

	fmt.Println()
}
