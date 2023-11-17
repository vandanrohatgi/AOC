package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func sol2() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	r := bytes.NewReader(f)
	reader := bufio.NewScanner(r)

	var totalRibbon int
	dimensions := make([]int, 3)
	for reader.Scan() {
		fmt.Fscanf(bytes.NewReader(reader.Bytes()), "%dx%dx%d", &dimensions[0], &dimensions[1], &dimensions[2])
		slices.Sort(dimensions)
		totalRibbon += 2 * (dimensions[0] + dimensions[1])
		totalRibbon += dimensions[0] * dimensions[1] * dimensions[2]
	}

	fmt.Println(totalRibbon)
}
