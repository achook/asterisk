package main

import (
	"bufio"
	"os"
	"testing"
)

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

func TestParseJSONIncorrect(t *testing.T) {
	// Parsable, valid JSON correct[N].json and parsable_incorrect[N].json located in test directory
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
