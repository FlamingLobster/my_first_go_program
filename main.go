package main

import (
	"bufio"
	"fmt"
	"koho/velocity"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if isSuccess, out := velocity.Allowed(scanner.Text()); isSuccess {
			fmt.Println(out)
		} else {
			fmt.Println("nope")
		}
	}
}
