package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func sol1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	priority := 0
	data := strings.Split(string(f), "\n")
	for _, i := range data[:len(data)-1] {
		sack := make(map[rune]interface{})
		var duplicate rune
		var score int
		for index, x := range i {
			_, ok := sack[x]
			if ok && index >= len(i)/2 {
				duplicate = x
				break
			}
			if index < len(i)/2 {
				sack[x] = nil
			}
		}
		if int(duplicate) >= 97 {
			score = int(duplicate) - 96
		} else {
			score = int(duplicate) - 38
		}
		priority += score
	}
	fmt.Println(priority)
}
