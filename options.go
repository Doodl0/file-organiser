package main

import (
	"fmt"
	"strings"
)

type Argument struct {
	variable    bool
	description string
}

var CopyFilesWithoutExtensions Argument = Argument{false, "Copy files without extensions? [Y/n]"}
var DeleteFilesAfterCopying Argument = Argument{false, "Delete original files after copying? [y/N]"}

func AskUserOption(option Argument) {

	// Ask user the option
	fmt.Print(option.description)
	// Get user's input
	var input string
	fmt.Scanln(&input)
	// Change option based on user input
	switch strings.ToLower(input) {
	case "y":
		option.variable = true

	case "n":
		option.variable = false
		fmt.Print(option)
	}
}
