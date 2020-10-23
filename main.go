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
		if err, output := velocity.Allowed(scanner.Text()); err != nil {
			fmt.Println(err)
		} else {
			println(output)
		}
	}
}
