package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func sol1() {
	var score = 0
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strats := strings.Split(string(f), "\n")
	for _, i := range strats {
		switch i {
		case "A X":
			score += 4 // A, X rock 1+3
		case "A Y":
			score += 8 // A rock, Y paper 2+6
		case "A Z":
			score += 3 // A rock, Z scissor 3+0
		case "B X":
			score += 1 // B paper, X rock 1+0
		case "B Y":
			score += 5 // B paper, Y paper 2+3
		case "B Z":
			score += 9 // B paper, Z scissor 3+6
		case "C X":
			score += 7 // C scissor, X rock 1+6
		case "C Y":
			score += 2 // C scissor, Y paper 2+0
		case "C Z":
			score += 6 // C scissor, Z scissor 3+3
		default:
			fmt.Printf("Did not find a match for %v", i)
		}
	}
	fmt.Println(score)
}
