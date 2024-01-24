package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vandanrohatgi/aoc/utility"
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

func pickNum(perms [][]int) [][]int {
	var result [][]int
	for a := 0; a <= len(perms[0]); a++ {
		for index, j := range perms {
			if sum(j) == EGGNOG {
				result = append(result, j)
			}
			perms[index] = j[:len(j)-1]

		}
		fmt.Printf("\n------------------ %+v ----------\n", perms)
	}
	return result
}

func sol1(f []string) {
	var conts = make([]int, len(f))
	var err error
	for i := range f {
		conts[i], err = strconv.Atoi(f[i])
		if err != nil {
			panic(err)
		}
	}
	perms := utility.Permutate(conts)
	fmt.Println(pickNum(perms))
}

func sol2(f []string) {

}
