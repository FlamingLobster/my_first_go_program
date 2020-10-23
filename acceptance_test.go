package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"koho/velocity"
	"os"
	"testing"
)

func TestAcceptable(t *testing.T) {
	allowed := make(map[int]string)
	setupResults(allowed)

	if inputFile, err := os.Open("input.txt"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()
			if loadFund, err := unmarhshalFunds(line); err != nil {
				continue
			} else {
				actuallyAllowed, actualOutput := velocity.Allowed(line)
				expectedOutput, expectedAllowed := allowed[loadFund.Id]
				if !actuallyAllowed && expectedAllowed {
					t.Error("Blocked when should have allowed")
				}
				if actuallyAllowed && !expectedAllowed {
					t.Error("Allowed when should have blocked")
				}
				if actuallyAllowed && expectedAllowed && actualOutput != expectedOutput {
					t.Error("Correctly allowed funds but output did not match expected output")
				}
			}
		}
	}
}

func unmarhshalFunds(line string) (*velocity.LoadFund, error) {
	var loadFund velocity.LoadFund
	if err := json.Unmarshal([]byte(line), &loadFund); err != nil {
		return nil, err
	} else {
		return &loadFund, nil
	}
}

func setupResults(allowed map[int]string) {
	resultFile, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(resultFile)
	var response velocity.Response
	for scanner.Scan() {
		if err := json.Unmarshal([]byte(scanner.Text()), &response); err != nil {
			fmt.Println("Failed to setup expected results")
			panic(err)
		} else {
			allowed[response.Id] = scanner.Text()
		}
	}
}
