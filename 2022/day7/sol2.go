package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func sol2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	sc.Scan()
	var fileDirectory = &fs{size: 0, fsTree: make(map[string]*fs)}
	recurseDir(fileDirectory, sc)
	var unusedSpace = 70000000 - fileDirectory.size
	var deleteDir = 70000000
	var requiredSpace = 30000000 - unusedSpace
	fmt.Println(unusedSpace, requiredSpace, fileDirectory.size)
	deleteDir = freeMem(fileDirectory, requiredSpace, deleteDir)
	fmt.Println(deleteDir)
}

func freeMem(dir *fs, requiredSpace, deleteDir int) int {
	if dir.size >= requiredSpace && dir.size <= deleteDir {
		deleteDir = dir.size
	}
	for _, v := range dir.fsTree {
		deleteDir = freeMem(v, requiredSpace, deleteDir)
	}
	return deleteDir
}
