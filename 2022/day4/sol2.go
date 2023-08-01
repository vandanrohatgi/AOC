package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func sol2() {
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
		// if (a <= c && a <= d && b == c && b <= d) || // I am so ashamed of this. I've let my family down, my fellow programmers and I feel in a way.... the whole world ;,)
		// (a <= c && a <= d && b >= c && b <= d) ||
		// (a == c && a <= d && b >= c && b == d) ||
		// (a >= c && a <= d && b >= c && b <= d) ||
		// (a >= c && a == d && b >= c && b >= d) {
		if b >= c && a <= d {
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}
