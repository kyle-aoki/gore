package model

import (
	"fmt"
	"gore/input"
	"gore/util"
	"log"
	"strconv"
)

var majorSpellingVariants = []string{"major", "1", "maj", "Major"}
var minorSpellingVariants = []string{"minor", "2", "min", "Minor"}
var patchSpellingVariants = []string{"patch", "3", "pat", "Patch"}

func GetNewRelease(release Release) Release {


	fmt.Printf("\nLatest release: %s", util.Blue(release.QualifiedName())+"\n\n")

	fmt.Println("Choose version increment for new release:")
	fmt.Println("(1) Major")
	fmt.Println("(2) Minor")
	fmt.Println("(3) Patch")

	inputText := input.WaitForInputText()

	newRelease := release

	if util.IsIn(inputText, majorSpellingVariants) {
		newRelease.Major = IncrementIntegerString(release.Major)
	} else if util.IsIn(inputText, minorSpellingVariants) {
		newRelease.Minor = IncrementIntegerString(release.Minor)
	} else if util.IsIn(inputText, patchSpellingVariants) {
		newRelease.Patch = IncrementIntegerString(release.Patch)
	} else {
		log.Fatal("Invalid input text.")
	}

	return newRelease
}

func IncrementIntegerString(integerString string) string {
	integer, err := strconv.ParseInt(integerString, 10, 10)
	if err != nil {
		log.Fatalf("Error with parsing string: %+v", err)
	}
	integer += 1
	return strconv.Itoa(int(integer))
}
