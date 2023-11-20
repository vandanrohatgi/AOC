package main

import (
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

	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

func coordinatesToString(x, y int) string {
	return strconv.Itoa(x) + ":" + strconv.Itoa(y)
}

func sol1(f []byte) {
	houseMap := make(map[string]int)
	var x, y int
	houseMap[coordinatesToString(x, y)] = 1
	for _, i := range f {
		switch rune(i) {
		case '^':
			y += 1
			houseMap[coordinatesToString(x, y)] += 1
		case 'v':
			y -= 1
			houseMap[coordinatesToString(x, y)] += 1
		case '<':
			x -= 1
			houseMap[coordinatesToString(x, y)] += 1
		case '>':
			x += 1
			houseMap[coordinatesToString(x, y)] += 1
		}
	}

	fmt.Print(len(houseMap))
}

func sol2(f []byte) {
	houseMap := make(map[string]int)
	var santa, robo = []int{0, 0}, []int{0, 0} // x,y coordinates
	var turnPtr = &santa                       // start with santa in the lead

	houseMap[coordinatesToString(santa[0], santa[1])] = 1 // starting house

	for _, i := range f {
		switch rune(i) {
		case '^':
			(*turnPtr)[1] += 1
			houseMap[coordinatesToString((*turnPtr)[0], (*turnPtr)[1])] += 1
		case 'v':
			(*turnPtr)[1] -= 1
			houseMap[coordinatesToString((*turnPtr)[0], (*turnPtr)[1])] += 1
		case '<':
			(*turnPtr)[0] -= 1
			houseMap[coordinatesToString((*turnPtr)[0], (*turnPtr)[1])] += 1
		case '>':
			(*turnPtr)[0] += 1
			houseMap[coordinatesToString((*turnPtr)[0], (*turnPtr)[1])] += 1
		}

		// change turns every character
		if turnPtr == &santa {
			turnPtr = &robo
		} else {
			turnPtr = &santa
		}
	}

	fmt.Print(len(houseMap))
}
