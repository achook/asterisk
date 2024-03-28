package main

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func parseJSON(reader *bufio.Reader) (RolePolicy, error) {
	var rolePolicy RolePolicy
	err := json.NewDecoder(reader).Decode(&rolePolicy)
	if err != nil {
		return RolePolicy{}, fmt.Errorf("Error decoding JSON: %s", err)
	}

	return rolePolicy, nil
}
