package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
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

type santaPass []rune

func (s *santaPass) setPass(input string) {
	*s = []rune(input)
}

func (s *santaPass) getString() string {
	return string(*s)
}

func (s *santaPass) incrementByOne(i int) {
	if (*s)[i] < 122 && (*s)[i] >= 97 {
		(*s)[i] += 1
	} else if (*s)[i] == 'z' && i != 0 {
		(*s)[i] = 'a'
		s.incrementByOne(i - 1)
	} // i==0 no action in case we can't go any further left
}

func sol1(f []byte) {
	currentPass := string(f)

	currentPass = generateNextPass(currentPass)
	fmt.Println(currentPass)
}

func generateNextPass(currentPass string) string {
	var s santaPass
	s.setPass(currentPass)
	s.incrementByOne(len(currentPass) - 1)

	currentPass = s.getString()

	for !hasTwoPairs(currentPass) || !notHaveBadLetters(currentPass) || !hasThreeLetters(currentPass) {
		s.incrementByOne(len(currentPass) - 1)
		currentPass = s.getString()
	}
	return currentPass
}

func hasTwoPairs(s string) bool {
	var count int
	var alreadyMatched byte
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] && s[i] != alreadyMatched {
			alreadyMatched = s[i]
			count += 1
		}
	}
	return count >= 2
}

func notHaveBadLetters(s string) bool {
	for _, i := range s {
		if i == 'i' || i == 'o' || i == 'l' {
			return false
		}
	}
	return true
}

func hasThreeLetters(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
	}
	return false
}

func sol2(f []byte) {
	currentPass := string(f)
	currentPass = generateNextPass(currentPass)
	currentPass = generateNextPass(currentPass)
	fmt.Println(currentPass)
}
