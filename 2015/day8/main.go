package main

import (
	"bufio"
	"bytes"
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
	f = bytes.TrimSpace(f)
	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

func sol1(f []byte) {
	r := bytes.NewReader(f)
	buf := bufio.NewScanner(r)
	var sum int
	for buf.Scan() {
		tmp, _ := strconv.Unquote(buf.Text())
		sum += len(buf.Text()) - len(tmp)
	}
	fmt.Println(sum)
}

func sol2(f []byte) {

}
