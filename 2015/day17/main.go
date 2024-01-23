package main

import (
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

	fs := string(f)
	fs = strings.TrimSpace(fs)
	stringInput := strings.Split(fs, "\n")
	if solution == 1 {
		sol1(stringInput)
	} else {
		sol2(stringInput)
	}
}

const EGGNOG = 25

func sum(input []int) int {
	var s int
	for _, i := range input {
		s += i
	}
	return s
}

type circularArray []int

var result2 [][]int

func (ca *circularArray) getElement(index int) int {
	if index >= len(*ca) {
		return (*ca)[index%len(*ca)]
	}
	return (*ca)[index]
}

func (ca *circularArray) getRange(i, j int) []int {
	var tmp []int
	for lol := i; lol <= j; lol++ {
		tmp = append(tmp, ca.getElement(lol))
	}
	return tmp
}

func Combine(cont circularArray) {
	for i := range cont {
		for j := i; j < len(cont)+i; j++ {
			result2 = append(result2, cont.getRange(i, j))
		}
	}
}

func sol1(f []string) {
	var conts = make(circularArray, len(f))
	var err error
	for i := range f {
		conts[i], err = strconv.Atoi(f[i])
		if err != nil {
			panic(err)
		}
	}
	// fmt.Println(Permutate(conts))
	// for _, i := range Permutate(conts) {
	// 	if sum(i) == EGGNOG {
	// 		count += 1
	// 	}
	// }
	Combine(conts)
	for _, i := range result2 {
		if sum(i) == EGGNOG {
			fmt.Println(i)
		}
	}
	// fmt.Println(result2)
}

func sol2(f []string) {

}
