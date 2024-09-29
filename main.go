package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("fileorganiser 1.0")
	files, err := os.ReadDir(".")

	// Confirm options with user
	AskUserOptions()

	//Loop through every file, and copy it to respective folder
	if err == nil {
		for _, file := range files {
			if !file.IsDir() {
				CopyFile(file.Name())
			}
		}
	}
}
