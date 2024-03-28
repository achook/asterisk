package main

import "testing"

func TestIfFileExistsCorrect(t *testing.T) {
	correctFilenames := []string{"utils.go", "utils_test.go", "test/correct1.json", "test/correct2.json"}

	for _, filename := range correctFilenames {
		if !checkIfFileExists(filename) {
			t.Errorf("File %s should exist", filename)
		}
	}
}

func TestIfFileExistsIncorrect(t *testing.T) {
	incorrectFilenames := []string{"nonexistent.go", "test/nonexistent.json"}

	for _, filename := range incorrectFilenames {
		if checkIfFileExists(filename) {
			t.Errorf("File %s should not exist", filename)
		}
	}
}
