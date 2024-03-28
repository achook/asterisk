package main

import (
	"bufio"
	"log"
)

// RolePolicy is a struct that represents the policy for a role.
// It is assumed to be defined elsewhere in the code.

// verifyAsterisk checks if all resources in the policy statements are "*".
// It returns false if all resources are "*", true otherwise.
func verifyAsterisk(policy RolePolicy) bool {
	if policy.PolicyDocument.Statement == nil {
		return true
	}

	if len(policy.PolicyDocument.Statement) == 0 {
		return true
	}

	for _, statement := range policy.PolicyDocument.Statement {
		if statement.Resource != "*" {
			return true
		}
	}

	return false
}

// CheckAsterisk takes a bufio.Reader and attempts to decode the JSON input into a RolePolicy struct.
// It then verifies if all resources in the policy statements are "*".
// It returns false if all resources are "*", true otherwise.
func CheckAsterisk(JSONReader *bufio.Reader) bool {
	rolePolicy, err := parseJSON(JSONReader)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return verifyAsterisk(rolePolicy)
}
