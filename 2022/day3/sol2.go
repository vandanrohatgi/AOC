package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func sol2() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	priority := 0
	data := strings.Split(string(f), "\n")
	for i := 0; i <= len(data)-3; i += 3 {
		var duplicate rune
		var score int
		duplicate = commonRune(data[i], data[i+1], data[i+2])
		if int(duplicate) >= 97 {
			score = int(duplicate) - 96
		} else {
			score = int(duplicate) - 38
		}
		priority += score
	}
	fmt.Println(priority)
}

func commonRune(s ...string) rune {
	commonMap := make(map[rune]int)
	for indexOuter, i := range s {
		for _, character := range i {
			if indexOuter == 0 {
				commonMap[character] = 0
			}
			val, ok := commonMap[character]
			if ok && val == indexOuter-1 {
				commonMap[character] = indexOuter
			}
		}
	}
	var commonKey rune
	for k, v := range commonMap {
		if v == len(s)-1 {
			commonKey = k
			break
		}
	}
	return commonKey
}
