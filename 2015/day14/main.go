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
	var bestDeerDist int
	var speed, limit, rest int
	var deer string
	for _, i := range strings.Split(string(f), "\n") {
		var Totaltime = 2503
		fmt.Sscanf(i, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &deer, &speed, &limit, &rest)
		singleRunTime := rest + limit
		wholeRun := Totaltime / singleRunTime
		Totaltime = Totaltime % singleRunTime
		if Totaltime > limit {
			Totaltime = limit
		}
		totalRun := Totaltime + wholeRun*limit
		TotalDist := speed * totalRun
		if TotalDist > bestDeerDist {
			bestDeerDist = TotalDist
		}
	}
	fmt.Println(bestDeerDist)
}

func sol2(f []byte) {

}
