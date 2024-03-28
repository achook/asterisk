package main

import (
	"bufio"
	"log"
)

// RolePolicy is a struct that represents the policy for a role.
// It is assumed to be defined elsewhere in the code.

// verifyAsterisk checks if all resources in the policy statements are "*".
// It returns true if all resources are "*", false otherwise.
func verifyAsterisk(policy RolePolicy) bool {
	if policy.PolicyDocument.Statement == nil {
		return false
	}

	if len(policy.PolicyDocument.Statement) == 0 {
		return false
	}

	for _, statement := range policy.PolicyDocument.Statement {
		if statement.Resource != "*" {
			return false
		}
	}

	return true
}

// CheckAsterisk takes a bufio.Reader and attempts to decode the JSON input into a RolePolicy struct.
// It then verifies if all resources in the policy statements are "*".
// It returns true if all resources are "*", false otherwise.
func CheckAsterisk(JSONReader *bufio.Reader) bool {
	rolePolicy, err := parseJSON(JSONReader)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return verifyAsterisk(rolePolicy)
}
