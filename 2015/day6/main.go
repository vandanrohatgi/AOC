package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

func sol1(f []byte) {

	instructions := strings.Split(string(f), "\n")

	//turn on 0,0 through 999,999
	//toggle 0,0 through 999,0
	//turn off 499,499 through 500,500
	parser := regexp.MustCompile(`((turn|toggle)\s?(on|off)?)\s(\d+,\d+\sthrough\s\d+,\d+)`)
	var cmd string
	var x1, y1, x2, y2, lit int
	var lights [1000][1000]bool

	for _, line := range instructions {
		matches := parser.FindStringSubmatch(line)
		cmd = matches[1] // submatch Group 1
		fmt.Fscanf(strings.NewReader(matches[4]), "%d,%d through %d,%d", &x1, &y1, &x2, &y2)

		swapIfWrong(&x1, &x2)
		swapIfWrong(&y1, &y2)

		switch cmd {
		case "turn on":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = true
				}
			}
		case "turn off":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = false
				}
			}

		case "toggle":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] = !lights[i][j]
				}
			}
		}
	}

	// we can't count the number of lit lights without this loop. if we try to
	// do it in the same loop above using a formlua like
	// lit+=(x2+1-x1)*(y2+1-y1) because that way we will be counting and turning
	// off lights which are already off giving as more than a million lit lights
	// at the end.

	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				lit += 1
			}
		}
	}

	fmt.Println(lit)
}

// a should be smaller than b
func swapIfWrong(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func sol2(f []byte) {
	instructions := strings.Split(string(f), "\n")

	//turn on 0,0 through 999,999
	//toggle 0,0 through 999,0
	//turn off 499,499 through 500,500
	parser := regexp.MustCompile(`((turn|toggle)\s?(on|off)?)\s(\d+,\d+\sthrough\s\d+,\d+)`)
	var cmd string
	var x1, y1, x2, y2, totalBright int
	var lights [1000][1000]int

	for _, line := range instructions {
		matches := parser.FindStringSubmatch(line)
		cmd = matches[1] // submatch Group 1
		fmt.Fscanf(strings.NewReader(matches[4]), "%d,%d through %d,%d", &x1, &y1, &x2, &y2)

		swapIfWrong(&x1, &x2)
		swapIfWrong(&y1, &y2)

		switch cmd {
		case "turn on":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] += 1
				}
			}
		case "turn off":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if lights[i][j] != 0 { // I thought using uint would let me escape this check,
						lights[i][j] -= 1 // but boy was I wrong.
					}
				}
			}

		case "toggle":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					lights[i][j] += 2
				}
			}
		}
	}

	for i := range lights {
		for j := range lights[i] {
			totalBright += lights[i][j]
		}
	}

	fmt.Println(totalBright)

}
