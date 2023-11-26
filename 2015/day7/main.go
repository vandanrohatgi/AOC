package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

// keep in mind the example may not be the only psoobile pattern in the input
// patterns of instructions
var directOp = regexp.MustCompile(`^(\d+|\w+) -> (\w+)$`)
var complementOp = regexp.MustCompile(`^NOT (\d+|\w+) -> (\w+)$`)
var binaryOp = regexp.MustCompile(`^(\w+) (\w+) (\d+|\w+) -> (\w+)$`)

func sol1(f []byte) {
	buf := bufio.NewScanner(bytes.NewReader(f))

	var register = make(map[string]uint16) // uint becuase apparently the left most bit is for sign of the integer making negative numbers
	for buf.Scan() {
		if direct.MatchString(buf.Text()) {

			directMatch := direct.FindStringSubmatch(buf.Text())
			fmt.Println(directMatch)
			var src interface{}
			src, err := strconv.Atoi(directMatch[1])
			dst := directMatch[2]
			if err != nil {
				register[dst] = register[src.(string)]
			}
			register[dst] = uint16(src.(int))

		} else if complement.MatchString(buf.Text()) {

			complimentMatch := complement.FindStringSubmatch(buf.Text())
			fmt.Println(complimentMatch)
			dst := complimentMatch[2]
			src := complimentMatch[1]
			register[dst] = ^register[src]

		} else {
            
            

			shiftMatch := shift.FindStringSubmatch(buf.Text())
			fmt.Println(shiftMatch)
			src := shiftMatch[1]
			shiftType := shiftMatch[2]
			shiftBy, err := strconv.Atoi(shiftMatch[3])
			if err != nil {
				log.Fatal(err)
			}
			dst := shiftMatch[4]

			if shiftType == "LSHIFT" {
				register[dst] = register[src] << shiftBy
			} else {
				register[dst] = register[src] >> shiftBy
			}

		} else if logic.MatchString(buf.Text()) {
			logicMatch := logic.FindStringSubmatch(buf.Text())
			fmt.Println(logicMatch)
			src1 := logicMatch[1]
			operator := logicMatch[2]
			src2 := logicMatch[3]
			dst := logicMatch[4]

			if operator == "AND" {
				register[dst] = register[src1] & register[src2]
			} else {
				register[dst] = register[src1] | register[src2]
			}
		}
	}

	fmt.Print(register["a"])

}

func sol2(f []byte) {

}
