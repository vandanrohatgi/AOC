package main

import (
	"fmt"
	"log"
	"os"
)

func sol2() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading input file")
	}

	var floor int
	fileData := string(f)

loopBreaker:
	for charNumber, char := range fileData {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
			if floor == -1 {
				fmt.Println(charNumber + 1) // +1 becuase index start at 0
				break loopBreaker
			}
		default:
		}
	}
}
