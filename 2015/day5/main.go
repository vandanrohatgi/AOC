package main

import (
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

	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

// PART 1
func passesAllChecks(s string) bool {
	vowels := map[rune]interface{}{
		'a': nil,
		'e': nil,
		'i': nil,
		'o': nil,
		'u': nil,
	}

	var vowelCount int
	var doubleLetter bool
	for i := 0; i < len(s)-1; i++ {
		subString := s[i : i+2]
		// check bad sub-strings first to fail early
		if subString == "ab" || subString == "pq" || subString == "cd" || subString == "xy" {
			return false
		}

		// check double letters
		if s[i] == s[i+1] {
			doubleLetter = true
		}
	}
	// check vowel
	for _, i := range s {
		if _, ok := vowels[i]; ok {
			vowelCount++
		}
	}

	return doubleLetter && vowelCount >= 3

}

func sol1(f []byte) {
	f = bytes.TrimSpace(f)
	cases := string(f)
	var niceStrings int
	for _, i := range strings.Split(cases, "\n") {
		if nice := passesAllChecks(i); nice {
			niceStrings++
			fmt.Println(i)
		}
	}
	fmt.Println(niceStrings)
}

// PART 2
func passesAllChecks2(s string) bool {
	return hasDoubleLetters(s) && hasSingleLetter(s)
}

// a b c d e f
// i=0, i+2 = 2
// j=i+2=2, j+2=4
func hasDoubleLetters(s string) bool {
	for i := 0; i <= len(s)-4; i++ {
		sub1 := s[i : i+2]
		for j := i + 2; j <= len(s)-2; j++ {
			sub2 := s[j : j+2]
			if sub1 == sub2 {
				return true
			}
		}
	}
	return false
}

func hasSingleLetter(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func sol2(f []byte) {
	f = bytes.TrimSpace(f)
	cases := string(f)
	var niceStrings int
	for _, i := range strings.Split(cases, "\n") {
		if nice := passesAllChecks2(i); nice {
			niceStrings++
			fmt.Println(i)
		}
	}
	fmt.Println(niceStrings)
}
