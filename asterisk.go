package main

import (
	"bufio"
	"log"
)

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

func CheckAsterisk(JSONReader *bufio.Reader) bool {
	rolePolicy, err := parseJSON(JSONReader)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return verifyAsterisk(rolePolicy)
}
