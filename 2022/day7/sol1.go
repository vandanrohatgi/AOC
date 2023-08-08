package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var memSum int

func sol1() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	sc.Scan()
	var fileDirectory = &fs{size: 0, fsTree: make(map[string]*fs)}
	recurseDir(fileDirectory, sc)
	traverseFs(fileDirectory)
	fmt.Println(memSum)
}

func traverseFs(dir *fs) {
	if dir.size <= 100000 {
		memSum += dir.size
	}
	for _, v := range dir.fsTree {
		traverseFs(v)
	}
	fmt.Println(dir)
}
