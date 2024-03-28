package main

import (
	"bufio"
	"fmt"
	"os"
)

// getCommandLineFilename retrieves the filename from the command line arguments.
// It returns an error if no filename is provided.
func getCommandLineFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no filename provided")
	}

	return os.Args[1], nil
}

// checkIfFileExists checks if a file with the given filename exists.
// It returns true if the file exists, false otherwise.
func checkIfFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// openFile opens a file with the given filename and returns a bufio.Reader for the file.
// It returns an error if the file cannot be opened.
func openFile(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	return bufio.NewReader(file), nil
}
