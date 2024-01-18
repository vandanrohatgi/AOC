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
	gift, _ := os.ReadFile("gift.txt")
	giftProps := make(map[string]int)
	auntsList := make(map[int]map[string]int)
	var tmpProp1, tmpProp2, tmpProp3 string
	var aunt, tmpAmount1, tmpAmount2, tmpAmount3 int
	for _, i := range strings.Split(strings.TrimSpace(string(gift)), "\n") {
		fmt.Sscanf(i, "%s %d", &tmpProp1, &tmpAmount1)
		giftProps[tmpProp1] = tmpAmount1
	}

	for _, i := range strings.Split(string(f), "\n") {
		fmt.Sscanf(i, "Sue %d: %s %d, %s %d, %s %d", &aunt, &tmpProp1, &tmpAmount1, &tmpProp2, &tmpAmount2, &tmpProp3, &tmpAmount3)
		auntsList[aunt] = make(map[string]int)
		auntsList[aunt][tmpProp1] = tmpAmount1 // can probably make this better with a cutom bufio scanner
		auntsList[aunt][tmpProp2] = tmpAmount2
		auntsList[aunt][tmpProp3] = tmpAmount3
	}

	var count int
	for auntNum, auntProps := range auntsList {
		count = 0
		for giftProp, giftPropAmount := range giftProps {
			if x, ok := auntProps[giftProp]; ok && x == giftPropAmount {
				count += 1
			}
		}
		if count >= 3 {
			fmt.Println(auntNum)
			break
		}
	}

}

func sol2(f []byte) {

}
