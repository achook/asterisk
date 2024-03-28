package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename, err := getCommandLineFilename()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if !checkIfFileExists(filename) {
		fmt.Printf("Error: file %s does not exist\n", filename)
		os.Exit(1)
	}

	reader, err := openFile(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	result := CheckAsterisk(reader)

	if result {
		log.Println("True")
	} else {
		log.Println("False")
	}
}
