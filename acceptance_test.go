package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"koho/velocity"
	"os"
	"testing"
)

/**
This test evolved from the beginning when I didn't fully understand all the requirements. As such, it's much too complicated.
This test should be simplified to simply traverse the two files in sequence to match the input to output. Although there is
a small complication with duplicated id and customer id not giving any output.

There's probably some way to construct a parameterized test with file sources but I've not looked into this.
*/
func TestAll(t *testing.T) {
	allowed := make(map[string]bool)
	setupResults(allowed)

	if inputFile, err := os.Open("input.txt"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()

			err, action, actualOutput := Allowed(line)
			if action == velocity.Ignore {
				continue
			}
			if err != nil {
				t.Error(err)
			}

			key := actualOutput
			_, present := allowed[key]
			if !present {
				t.Error(
					"Incorrect output\n",
					"Input:    "+line,
					"\n",
					"Actual:   "+actualOutput,
				)
			}
			delete(allowed, key)
		}
		if len(allowed) != 0 {
			t.Error("Failed to match all rows in expected output")
		}
	}
}

func unmarhshalFunds(line string) (*velocity.Funds, error) {
	var loadFund velocity.Funds
	if err := json.Unmarshal([]byte(line), &loadFund); err != nil {
		return nil, err
	} else {
		return &loadFund, nil
	}
}

func setupResults(allowed map[string]bool) {
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
			allowed[scanner.Text()] = true
		}
	}
}
