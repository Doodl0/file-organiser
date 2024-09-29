package main

import (
	"fmt"
)

type Argument struct {
	variable    bool
	description string
}

var CopyFilesWithoutExtensions Argument = Argument{true, "Copy files without extensions? [Y/n]"}
var DeleteFilesAfterCopying Argument = Argument{false, "Delete original files after copying? [y/N]"}

func AskUserOptions() {

	// Ask user the option
	fmt.Print(CopyFilesWithoutExtensions.description)
	// Get user's input
	var input string
	fmt.Scanln(&input)
	// Change option based on user input
	if input == "y" || input == "Y" {
		CopyFilesWithoutExtensions.variable = true
	} else if input == "n" || input == "N" {
		CopyFilesWithoutExtensions.variable = false
	}

	// Ask user the option
	fmt.Print(DeleteFilesAfterCopying.description)
	// Get user's input
	fmt.Scanln(&input)
	// Change option based on user input
	if input == "y" || input == "Y" {
		DeleteFilesAfterCopying.variable = true
	} else if input == "n" || input == "N" {
		DeleteFilesAfterCopying.variable = false
	}
}
