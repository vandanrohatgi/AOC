package main

import (
	"bytes"
	"encoding/json"
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

func sol1(f []byte) {
	var m interface{}
	json.Unmarshal(f, &m)
	countNums(m)
	fmt.Println(sum)
}

var sum float64

// I'm disgusted by the amount of times I have to write interface{}
func countNums(m interface{}) {
	if val, intOk := m.(float64); intOk {
		sum += val
		return
	}

	switch input := m.(type) {
	case map[string]interface{}:
		if hasRed(input) {
			return
		}
		for _, v := range input {
			countNums(v)
		}
	case []interface{}:
		for _, i := range input {
			countNums(i)
		}
	}
}

func hasRed(m map[string]interface{}) bool {
	for _, v := range m {
		if v == "red" {
			return true
		}
	}
	return false
}

func sol2(f []byte) {
	var m interface{}
	json.Unmarshal(f, &m)
	countNums(m)
	fmt.Println(sum)
}
