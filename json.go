package main

import (
	"bufio"
	"encoding/json"
	"fmt"
)

// RolePolicy is a struct that represents the policy for a role.
// It is assumed to be defined elsewhere in the code.

// parseJSON takes a bufio.Reader and attempts to decode the JSON input into a RolePolicy struct.
// It returns the RolePolicy and any error encountered.
func parseJSON(reader *bufio.Reader) (RolePolicy, error) {
	var rolePolicy RolePolicy
	err := json.NewDecoder(reader).Decode(&rolePolicy)
	if err != nil {
		return RolePolicy{}, fmt.Errorf("Error decoding JSON: %s", err)
	}

	return rolePolicy, nil
}
