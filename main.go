package main

import (
	"bufio"
	"fmt"
	"koho/velocity"
	"os"
)

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
		err, action, output := velocity.Allowed(scanner.Text())
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

func write(file *os.File, s string) {
	_, err := file.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}
