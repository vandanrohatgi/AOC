package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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

const EGGNOG = 150

func sum(input []int) int {
	var s int
	for _, i := range input {
		s += i
	}
	return s
}

func All(set []int) (subsets [][]int) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		if sum(subset) == EGGNOG {
			subsets = append(subsets, subset)
		}
	}
	return subsets
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
	var count int
	combs := All(conts)
	for _, i := range combs {
		if sum(i) == EGGNOG {
			count += 1
		}
	}
	fmt.Println(count)
}

func sol2(f []string) {
	var conts = make([]int, len(f))
	var err error
	for i := range f {
		conts[i], err = strconv.Atoi(f[i])
		if err != nil {
			panic(err)
		}
	}
	var count int
	var mini = math.MaxInt
	combs := All(conts)

	for _, i := range combs {
		if len(i) < mini {
			mini = len(i)
			count = 1
		} else if len(i) == mini {
			count += 1
		}
	}

	fmt.Println(count)
}
