package testrunner

import (
	"encoding/json"
	"os"
	"strings"
)

type Problem struct {
	ID        string     `json:"id"`
	TestCases []TestCase `json:"testCases"`
}

type TestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
}

func LoadData() (Problem, error) {
	f, err := os.Open("../data.json")
	if err != nil {
		return Problem{}, err
	}
	defer f.Close()

	problem := Problem{}
	if err := json.NewDecoder(f).Decode(&problem); err != nil {
		return Problem{}, err
	}

	return problem, nil
}

// This is subject to change.

func CheckCode(language string, output string, /*code string*/) string {
	switch language {
	case "bash":
		return verdict("Hello world!", output)
	case "python":
		return verdict("Hello world!", output)
	case "go":
		return verdict("Hello world!", output)
	default:
		return "REJECTED"
	}
}

func verdict(expected, actual string) string {
	if strings.TrimSpace(expected) == strings.TrimSpace(actual) {
		return "ACCEPTED"
	}
	return "REJECTED"
}