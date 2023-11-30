package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
		sol1(bytes.TrimSpace(f))
	} else {
		sol2(bytes.TrimSpace(f))
	}
}

var instructions = make(map[string][]string)
var register = make(map[string]uint16)

func sol1(f []byte) {
	ins := strings.Split(string(f), "\n")
	for _, i := range ins {
		parts := strings.Split(i, " -> ")
		instructions[parts[1]] = strings.Split(parts[0], " ")
	}
	fmt.Println(calculateInstruction("a"))
}

func calculateInstruction(wire string) uint16 {
	if val, numeric := isNumeric(wire); numeric {
		return val
	}
	if val, ok := register[wire]; ok {
		return val
	}
	instruction := instructions[wire]
	switch {
	case len(instruction) == 1:
		register[wire] = calculateInstruction(instruction[0])
	case instruction[0] == "NOT":
		register[wire] = ^calculateInstruction(instruction[1])
	case instruction[1] == "AND":
		register[wire] = calculateInstruction(instruction[0]) & calculateInstruction(instruction[2])
	case instruction[1] == "OR":
		register[wire] = calculateInstruction(instruction[0]) | calculateInstruction(instruction[2])
	case instruction[1] == "LSHIFT":
		register[wire] = calculateInstruction(instruction[0]) << calculateInstruction(instruction[2])
	case instruction[1] == "RSHIFT":
		register[wire] = calculateInstruction(instruction[0]) >> calculateInstruction(instruction[2])
	}
	return register[wire]
}

func isNumeric(s string) (uint16, bool) {
	intString, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	} else {
		return uint16(intString), true
	}
}

func sol2(f []byte) {
	ins := strings.Split(string(f), "\n")
	for _, i := range ins {
		parts := strings.Split(i, " -> ")
		instructions[parts[1]] = strings.Split(parts[0], " ")
	}
	tmp := calculateInstruction("a")
	instructions["b"] = []string{strconv.Itoa(int(tmp))}
	register = make(map[string]uint16)
	fmt.Println(calculateInstruction("a"))
}
