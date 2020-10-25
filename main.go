package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	. "koho/velocity"
	"os"
)

var limits = NewLimit()

func Reset() {
	limits = NewLimit()
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
		if action == Ignore {
			continue
		}
		write(outputFile, output)
	}
}

/**
Error is on the left because I'm more used to it. It seems that Go prefers the error on the right. Will have to refactor
this.
*/
func Allowed(event string) (error, int, string) {
	var loadFund Funds
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return err, Ignore, ""
	} else {
		action := limits.Allowed(loadFund)
		switch action {
		case Accept:
			if output, err := json.Marshal(Accepted(&loadFund)); err != nil {
				return err, Ignore, ""
			} else {
				return nil, Accept, string(output)
			}
		case Reject:
			if output, err := json.Marshal(Denied(&loadFund)); err != nil {
				return err, Ignore, ""
			} else {
				return nil, Reject, string(output)
			}
		case Ignore:
			return nil, Ignore, ""
		}
	}
	return errors.New("should not get here"), Ignore, ""
}

func write(file *os.File, s string) {
	_, err := file.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}
