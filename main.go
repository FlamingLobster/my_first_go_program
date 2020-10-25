package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"koho/velocity"
	"os"
)

var limits = velocity.NewLimit()

func Reset() {
	limits = velocity.NewLimit()
}

func main() {
	outputFile, err := os.Create("produced.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		err, action, output := Allowed(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		if action == velocity.Ignore {
			continue
		}
		write(outputFile, output)
	}
}

func Allowed(event string) (error, int, string) {
	var loadFund velocity.Funds
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return err, velocity.Ignore, ""
	} else {
		action := limits.Allowed(loadFund)
		switch action {
		case velocity.Accept:
			if output, err := json.Marshal(velocity.Accepted(&loadFund)); err != nil {
				return err, velocity.Ignore, ""
			} else {
				return nil, velocity.Accept, string(output)
			}
		case velocity.Reject:
			if output, err := json.Marshal(velocity.Denied(&loadFund)); err != nil {
				return err, velocity.Ignore, ""
			} else {
				return nil, velocity.Reject, string(output)
			}
		case velocity.Ignore:
			return nil, velocity.Ignore, ""
		}
	}
	return errors.New("should not get here"), velocity.Ignore, ""
}

func write(file *os.File, s string) {
	_, err := file.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}
