package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func sol1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataRegex := regexp.MustCompile(`(\d{1,999})-(\d{1,999}),(\d{1,999})-(\d{1,999})`)
	matches := dataRegex.FindAllStringSubmatch(string(f), -1)
	var overlaps int
	for _, i := range matches {
		a, _ := strconv.Atoi(i[1])
		b, _ := strconv.Atoi(i[2])
		c, _ := strconv.Atoi(i[3])
		d, _ := strconv.Atoi(i[4])
		if c >= a && d <= b {
			overlaps += 1
		} else if a >= c && b <= d {
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}
