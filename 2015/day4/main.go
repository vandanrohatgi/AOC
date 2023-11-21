package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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

func getHash(secret []byte, number int) string {
	secret = append(secret, []byte(strconv.Itoa(number))...)
	md5Sum := md5.Sum(secret)
	sum := hex.EncodeToString(md5Sum[:])
	return sum
}

func sol1(f []byte) {
	f = bytes.TrimSpace(f)
	for i := 0; ; i++ {
		sum := getHash(f, i)
		if sum[:5] == "00000" {
			fmt.Println(i)
			break
		}
	}
}

func sol2(f []byte) {
	f = bytes.TrimSpace(f)
	for i := 0; ; i++ {
		sum := getHash(f, i)
		if sum[:6] == "000000" {
			fmt.Println(i)
			break
		}
	}

}
