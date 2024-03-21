package asterisk

import (
	"bufio"
	"log"
)

func CheckAsterisk(JSONReader *bufio.Reader) bool {
	rolePolicy, err := parseJSON(JSONReader)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	if rolePolicy.PolicyDocument.Statement[0].Resource == "*" {
		return true
	}

	return false
}
