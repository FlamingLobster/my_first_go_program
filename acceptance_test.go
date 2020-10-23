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
	allowed := make(map[velocity.Tuple]string)
	setupResults(allowed)

	if inputFile, err := os.Open("input.txt"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()

			loadFund, err := unmarhshalFunds(line)
			if err != nil {
				t.Error("Could not unmarshal input json")
			}
			err, actualOutput := velocity.Allowed(line)
			if err != nil {
				t.Error(err)
			}

			expectedOutput, present := allowed[velocity.KeyOf(loadFund.Id, loadFund.CustomerId)]
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

func unmarhshalFunds(line string) (*velocity.LoadFund, error) {
	var loadFund velocity.LoadFund
	if err := json.Unmarshal([]byte(line), &loadFund); err != nil {
		return nil, err
	} else {
		return &loadFund, nil
	}
}

func setupResults(allowed map[velocity.Tuple]string) {
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
			allowed[velocity.KeyOf(response.Id, response.CustomerId)] = scanner.Text()
		}
	}
}
