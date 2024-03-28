package main

type RolePolicy struct {
	PolicyName     string `json:"PolicyName"`
	PolicyDocument struct {
		Version   string `json:"Version"`
		Statement []struct {
			Sid      string   `json:"Sid"`
			Effect   string   `json:"Effect"`
			Action   []string `json:"Action"`
			Resource string   `json:"Resource"`
		} `json:"Statement"`
	} `json:"PolicyDocument"`
}
