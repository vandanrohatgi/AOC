package main

import (
	"fmt"
	"log"
	"os"
)

func sol1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading input file")
	}

	var floor int
	fileData := string(f)
	for _, char := range fileData {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
		}

		fmt.Println(string(char), floor)
	}
	fmt.Println(floor)
}
