package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func sol1() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	r := bytes.NewReader(f)
	reader := bufio.NewScanner(r)

	var totalPaper int
	dimensions := make([]int, 3)
	for reader.Scan() {
		fmt.Fscanf(bytes.NewReader(reader.Bytes()), "%dx%dx%d", &dimensions[0], &dimensions[1], &dimensions[2])
		slices.Sort(dimensions)
		totalPaper += 2 * (dimensions[0]*dimensions[1] + dimensions[1]*dimensions[2] + dimensions[0]*dimensions[2])
		totalPaper += dimensions[0] * dimensions[1]
	}

	fmt.Println(totalPaper)
}
