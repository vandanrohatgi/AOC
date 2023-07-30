package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sol1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(f), "\n")
	var sum, tmpSum int
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			if tmpSum > sum {
				sum = tmpSum
			}
			tmpSum = 0
			continue
		}
		integer, _ := strconv.Atoi(data[i])
		tmpSum += integer
	}
	fmt.Println(sum)
}
