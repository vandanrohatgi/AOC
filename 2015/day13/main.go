package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var solution int
	var isTest bool
	var f []byte
	var err error

	flag.IntVar(&solution, "s", 1, "Solution to execute (1 or 2)")
	flag.BoolVar(&isTest, "t", false, "run solution in testing mode")
	flag.Parse()

	if isTest {
		f, err = os.ReadFile("test.txt")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err = os.ReadFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	f = bytes.TrimSpace(f)
	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func createPermutations(s []string) [][]string {
	if len(s) == 1 {
		return [][]string{s}
	}

	var result [][]string

	for i := 0; i < len(s); i++ {
		tmp := make([]string, len(s))
		copy(tmp, s)
		tmp = append(tmp[:i], tmp[i+1:]...)
		for _, k := range createPermutations(tmp) {
			result = append(result, append([]string{s[i]}, k...))
		}
	}
	return result
}

var happiness = make(map[string]map[string]int)

func sol1(f []byte) {

	buf := bufio.NewScanner(bytes.NewReader(f))
	var s1, s2, sign string
	var unit int
	for buf.Scan() {
		fmt.Sscanf(strings.Trim(buf.Text(), "."), "%s would %s %d happiness units by sitting next to %s", &s1, &sign, &unit, &s2)
		if _, ok := happiness[s1]; !ok {
			happiness[s1] = make(map[string]int)
		}
		if sign == "lose" {
			unit *= -1
		}
		happiness[s1][s2] = unit
	}

	var family []string
	for i := range happiness {
		family = append(family, i)
	}

	allPerms := createPermutations(family)

	var maxHappy int
	for _, i := range allPerms {
		tmp := calcHappy(i)
		if tmp > maxHappy {
			maxHappy = tmp
		}
	}
	fmt.Println(maxHappy)
}

func calcHappy(s []string) int {
	var happy int
	for i := range s {
		happy += happiness[s[i]][s[(i+1)%len(s)]]
		happy += happiness[s[(i+1)%len(s)]][s[i]]
	}
	return happy
}

func sol2(f []byte) {

}
