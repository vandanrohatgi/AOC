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

var chart = make(map[string][]int)

const TEASPOONS = 100

func calcScore(ing1, ing2, ing3, ing4 string, ing1cont, ing2cont, ing3cont, ing4cont int) int {
	capa := max(0, chart[ing1][0]*ing1cont+chart[ing2][0]*ing2cont+chart[ing3][0]*ing3cont+chart[ing4][0]*ing4cont)
	dura := max(0, chart[ing1][1]*ing1cont+chart[ing2][1]*ing2cont+chart[ing3][1]*ing3cont+chart[ing4][1]*ing4cont)
	flav := max(0, chart[ing1][2]*ing1cont+chart[ing2][2]*ing2cont+chart[ing3][2]*ing3cont+chart[ing4][2]*ing4cont)
	text := max(0, chart[ing1][3]*ing1cont+chart[ing2][3]*ing2cont+chart[ing3][3]*ing3cont+chart[ing4][3]*ing4cont)
	return capa * dura * flav * text
}

func sol1(f []byte) {
	var ingredient string
	var capacity, durability, flavor, texture, calories int
	for _, i := range strings.Split(string(f), "\n") {
		fmt.Sscanf(i, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ingredient, &capacity, &durability, &flavor, &texture, &calories)
		ingredient = strings.Trim(ingredient, ":")
		chart[ingredient] = []int{capacity, durability, flavor, texture}
	}
	var maxScore int
	for i := 0; i < TEASPOONS; i++ {
		for j := 0; j < TEASPOONS; j++ {
			for k := 0; k < TEASPOONS; k++ {
				for l := 0; l < TEASPOONS; l++ {
					if i+j+k+l > TEASPOONS {
						continue
					}
					if score := calcScore("Sprinkles", "PeanutButter", "Frosting", "Sugar", i, j, k, l); score > maxScore {
						maxScore = score
					}
				}
			}
		}
	}
	fmt.Println(maxScore)

}

func sol2(f []byte) {

}
