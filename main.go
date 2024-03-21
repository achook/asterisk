package main

import (
	"asterisk"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getCommandLineFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no filename provided")
	}

	return os.Args[1], nil
}

func checkIfFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// open file and return reader
func openFile(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	return bufio.NewReader(file), nil
}

func main() {
	filename, err := getCommandLineFilename()
	if err != nil {
		log.Fatalf("Error getting filename: %s", err)
	}

	if !checkIfFileExists(filename) {
		log.Fatalf("File does not exist: %s", filename)
	}

	reader, err := openFile(filename)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	result := asterisk.CheckAsterisk(reader)

	if result {
		log.Println("True")
	} else {
		log.Println("False")
	}
}
