package main

import (
	"bufio"
	"os"
	"testing"
)

// TestParseJSONCorrect tests the parseJSON function with JSON files that should be parsable.
func TestParseJSONCorrect(t *testing.T) {
	var correctFilename = []string{"test/correct1.json", "test/correct2.json", "test/parsable_incorrect1.json",
		"test/duplicate_resource.json", "test/missing_resource.json", "test/missing_other_field.json",
		"test/too_many_fields.json"}

	for _, filename := range correctFilename {
		t.Logf("Testing file: %s", filename)

		file, err := os.Open(filename)
		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}
		reader := bufio.NewReader(file)
		_, err = parseJSON(reader)
		if err != nil {
			t.Errorf("Error parsing JSON: %s", err)
		}
	}

}

// TestParseJSONIncorrect tests the parseJSON function with JSON files that should not be parsable.
func TestParseJSONIncorrect(t *testing.T) {
	var incorrectFilename = []string{"test/empty.json", "test/incorrect1.json", "test/incorrect2.json",
		"test/incorrect3.json"}

	for _, filename := range incorrectFilename {
		t.Logf("Testing file: %s", filename)

		file, err := os.Open(filename)
		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}
		reader := bufio.NewReader(file)
		_, err = parseJSON(reader)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	}
}

// TestCommandlineFilenameCorrect tests the getCommandLineFilename function with command line arguments that should be valid.
func TestCommandlineFilenameCorrect(t *testing.T) {
	os.Args = []string{"cmd", "filename", "filename2"}

	filename, err := getCommandLineFilename()
	if err != nil {
		t.Errorf("Error getting filename: %s", err)
	}
	if filename != "filename" {
		t.Errorf("Expected filename, got %s", filename)
	}
}

// TestCommandlineFilenameIncorrect tests the getCommandLineFilename function with command line arguments that should be invalid.
func TestCommandlineFilenameIncorrect(t *testing.T) {
	os.Args = []string{"cmd"}

	_, err := getCommandLineFilename()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
