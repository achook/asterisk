package main

import (
	"bufio"
	"fmt"
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
