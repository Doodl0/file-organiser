package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func CopyFile(filename string) {

	var directoryname string = "no extension"

	// Split filename into name and extension
	splitname := strings.Split(filename, ".")

	// If file has an extension, set the desired directory to that file extension
	// If above is false and we do not want to copy files without extensions, return
	if len(splitname) > 1 {
		directoryname = splitname[1]
	} else if !CopyFilesWithoutExtensions.variable {
		return
	}

	// If directory does not exist, make it
	if _, err := os.Stat(directoryname); errors.Is(err, os.ErrNotExist) {
		CreateFolder(filename)
	}

	// Read File
	file, err1 := os.ReadFile(filename)
	fmt.Println("Read file", filename)

	// Error handling
	errorCheck(err1)

	// Write new file
	err2 := os.WriteFile(directoryname+"/"+filename, file, 0777)
	fmt.Println("Wrote file", filename, "in directory", directoryname)

	// Error handling
	errorCheck(err2)

	// If we want to delete original files, delete it
	if DeleteFilesAfterCopying.variable {
		fmt.Println(DeleteFilesAfterCopying.variable)
		err3 := os.Remove(filename)

		// Error handling
		errorCheck(err3)
	}
}

// Create a folder for a file extension
func CreateFolder(filename string) {
	splitname := strings.Split(filename, ".")

	if len(splitname) > 1 {
		os.Mkdir(splitname[1], 0777)
		fmt.Println("Made directory", splitname[1])
	} else {
		os.Mkdir("no extension", 0777)
		fmt.Println("Made directory no extension")
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
