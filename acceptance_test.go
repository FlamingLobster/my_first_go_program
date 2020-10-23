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
				if err, actualOutput := velocity.Allowed(line); err != nil {
					t.Error(err)
				} else {
					expectedOutput, present := allowed[loadFund.Id]
					if !present {
						t.Error("No output produced when output is expected")
					}
					if actualOutput != expectedOutput {
						t.Error(
							"Incorrect output\n",
							"Expected: "+expectedOutput,
							"\n",
							"Actual:   "+actualOutput,
						)
					}
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
