package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sol2() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(f), "\n")
	top3 := []int{0, 0, 0}
	var tmpSum int
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			if tmpSum > top3[0] {
				top3[2] = top3[1]
				top3[1] = top3[0]
				top3[0] = tmpSum

			} else if tmpSum > top3[1] {
				top3[1] = top3[0]
				top3[1] = tmpSum
			} else if tmpSum > top3[2] {
				top3[2] = tmpSum
			}
			tmpSum = 0
			continue
		}
		integer, _ := strconv.Atoi(data[i])
		tmpSum += integer
	}
	fmt.Println(top3[0] + top3[1] + top3[2])

}
