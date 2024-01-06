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
		fmt.Println(TotalDist)
	}
	// fmt.Println(bestDeerDist)
}

const Totaltime = 2503

type Deer struct {
	speed         int
	limit         int
	rest          int
	distTravelled int
	points        int
	runTime       int
	restTime      int
}

func (d *Deer) increment() {
	if d.runTime < d.limit {
		d.distTravelled += d.speed
		d.runTime += 1
	} else if d.restTime < d.rest {
		d.restTime += 1
	} else {
		d.runTime = 1
		d.distTravelled += d.speed
		d.restTime = 0
	}
}

func (d *Deer) addPoint() {
	d.points += 1
}

func sol2(f []byte) {
	var speed, limit, rest int
	var deerName string
	var deers = make(map[string]*Deer)

	for _, i := range strings.Split(string(f), "\n") {
		_, err := fmt.Sscanf(i, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &deerName, &speed, &limit, &rest)
		if err != nil {
			panic(err)
		}
		deers[deerName] = &Deer{speed: speed, limit: limit, rest: rest, distTravelled: 0}
	}

	var currentBestDist int
	for i := 0; i < Totaltime; i++ {
		currentBestDist = 0
		for _, j := range deers {
			j.increment()
			if currentBestDist < j.distTravelled {
				currentBestDist = j.distTravelled
			}
		}

		// becuase we give point to each deer who is tied
		for _, j := range deers {
			if j.distTravelled == currentBestDist {
				j.addPoint()
			}
		}
	}

	for _, j := range deers {
		fmt.Println(j.points)
	}
}
