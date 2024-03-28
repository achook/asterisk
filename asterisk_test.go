package main

import (
	"encoding/json"
	"os"
	"testing"
)

// TestCheckAsteriskCorrect tests the verifyAsterisk function with JSON files that should return false.
func TestCheckAsteriskCorrect(t *testing.T) {
	correctFilenames := []string{"test/correct1.json", "test/correct2.json", "test/correct3.json"}

	for _, filename := range correctFilenames {
		t.Logf("Testing file: %s", filename)

		file, err := os.Open(filename)
		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}

		var rolePolicy RolePolicy
		err = json.NewDecoder(file).Decode(&rolePolicy)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		result := verifyAsterisk(rolePolicy)

		if result {
			t.Errorf("Expected true, got false")
		}

	}
}

// TestCheckAsteriskIncorrect tests the verifyAsterisk function with JSON files that should return true.
func TestCheckAsteriskIncorrect(t *testing.T) {
	incorrectFilenames := []string{"test/parsable_incorrect1.json", "test/parsable_incorrect2.json",
		"test/parsable_incorrect3.json", "test/parsable_incorrect4.json", "test/parsable_incorrect5.json",
		"test/parsable_incorrect6.json", "test/parsable_incorrect7.json",
		"test/duplicate_resource.json", "test/missing_resource.json"}

	for _, filename := range incorrectFilenames {
		t.Logf("Testing file: %s", filename)

		file, err := os.Open(filename)
		if err != nil {
			t.Errorf("Error opening file: %s", err)
		}

		var rolePolicy RolePolicy
		err = json.NewDecoder(file).Decode(&rolePolicy)
		if err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		result := verifyAsterisk(rolePolicy)

		if !result {
			t.Errorf("Expected false, got true")
		}
	}
}
