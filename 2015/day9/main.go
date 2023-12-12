package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
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

var graph = make(map[string]map[string]int)

func sol1(f []byte) {
	input := strings.Split(string(f), "\n")
	var src, dst string
	var cost int

	for _, i := range input {
		fmt.Fscanf(strings.NewReader(i), "%s to %s = %d", &src, &dst, &cost)
		if _, found := graph[src]; !found {
			graph[src] = make(map[string]int)
		}
		if _, found := graph[dst]; !found {
			graph[dst] = make(map[string]int)
		}
		graph[src][dst] = cost
		graph[dst][src] = cost
	}
	var nodes []string
	for k := range graph {
		nodes = append(nodes, k)
	}

	allCombinations := comb[string](nodes)

}

func comb[T any](i []T) [][]T {
	if len(i) == 1 {
		return [][]T{i}
	}

	var newCombs [][]T
	for k := 0; k < len(i); k++ {
		newI := make([]T, len(i))
		copy(newI, i)
		newI = append(newI[:k], newI[k+1:]...)

		for _, jk := range comb(newI) {
			jk = append([]T{i[k]}, jk...)
			newCombs = append(newCombs, jk)
		}
	}

	return newCombs
}

func sol2(f []byte) {

}
