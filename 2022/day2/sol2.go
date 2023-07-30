package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func sol2() {
	var score = 0
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strats := strings.Split(string(f), "\n")
	for _, i := range strats {
		switch i {
		case "A X":
			score += 3
		case "A Y":
			score += 4
		case "A Z":
			score += 8
		case "B X":
			score += 1
		case "B Y":
			score += 5
		case "B Z":
			score += 9
		case "C X":
			score += 2
		case "C Y":
			score += 6
		case "C Z":
			score += 7
		default:
			fmt.Printf("Did not find a match for %v", i)
		}
	}
	fmt.Println(score)
}
