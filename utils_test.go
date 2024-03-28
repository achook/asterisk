package main

import "testing"

// TestIfFileExistsCorrect tests the checkIfFileExists function with filenames that should exist.
func TestIfFileExistsCorrect(t *testing.T) {
	correctFilenames := []string{"utils.go", "utils_test.go", "test/correct1.json", "test/correct2.json"}

	for _, filename := range correctFilenames {
		if !checkIfFileExists(filename) {
			t.Errorf("File %s should exist", filename)
		}
	}
}

// TestIfFileExistsIncorrect tests the checkIfFileExists function with filenames that should not exist.
func TestIfFileExistsIncorrect(t *testing.T) {
	incorrectFilenames := []string{"nonexistent.go", "test/nonexistent.json"}

	for _, filename := range incorrectFilenames {
		if checkIfFileExists(filename) {
			t.Errorf("File %s should not exist", filename)
		}
	}
}

// TestOpenFileCorrect tests the openFile function with filenames that should be able to be opened.
func TestOpenFileCorrect(t *testing.T) {
	correctFilenames := []string{"utils.go", "utils_test.go", "test/correct1.json", "test/correct2.json"}

	for _, filename := range correctFilenames {
		reader, err := openFile(filename)
		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}
		if reader == nil {
			t.Errorf("Reader should not be nil")
		}
	}
}

// TestOpenFileIncorrect tests the openFile function with filenames that should not be able to be opened.
func TestOpenFileIncorrect(t *testing.T) {
	incorrectFilenames := []string{"nonexistent.go", "test/nonexistent.json"}

	for _, filename := range incorrectFilenames {
		reader, err := openFile(filename)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if reader != nil {
			t.Errorf("Reader should be nil")
		}
	}
}
