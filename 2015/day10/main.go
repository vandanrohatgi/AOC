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

	f = bytes.TrimSpace(f)
	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

func sol1(f []byte) {
	input := string(f)
	var output string

	for i := 0; i < 40; i++ {
		output = lookAndSay(input)
		input = output
	}
	fmt.Println(len(output))
}

func lookAndSay(input string) string {
	var output strings.Builder
	var count = 0
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 || input[i] != input[i+1] {
			output.WriteString(fmt.Sprintf("%d%s", count+1, string(input[i])))
			count = 0
		} else {
			count += 1
		}
	}

	return output.String()
}

func sol2(f []byte) {
	input := string(f)
	var output string

	for i := 0; i < 50; i++ {
		output = lookAndSay(input)
		input = output
	}
	fmt.Println(len(output))
}
