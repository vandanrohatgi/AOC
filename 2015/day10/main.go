package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

var memo map[string]string

func sol1(f []byte) {
	input := string(f)
	var output string

	memo = make(map[string]string)

	for i := 0; i < 40; i++ {
		output = lookAndSay(input)
		input = output
	}
	fmt.Println(len(output))
}

func getSequence(input string, iter int) string {
	if len(input) == 1 {
		return "1" + input
	}
	if val, ok := memo[input]; ok {
		return val
	}
    getSequence(input[], iter int)
}

func lookAndSay(input string) string {
	var output string
	var current = input[0]
	var count = 1
	for i := 1; i < len(input); i++ {
		if current == input[i] {
			count += 1
		} else {
			output = output + strconv.Itoa(count) + string(current)
			count = 1
		}
		current = input[i]
	}

	output = output + strconv.Itoa(count) + string(current) // for the last char(S) who didnt get their loops
	return output
}

// add a dictionary to keep a record of previous already calculated sequences
// since they will repeat.
func sol2(f []byte) {
	input := string(f)
	var output string
	for i := 0; i < 50; i++ {
		output = lookAndSay(input)
		input = output
	}
	fmt.Println(len(output))
}
