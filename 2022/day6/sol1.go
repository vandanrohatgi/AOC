package main

import (
	"fmt"
	"log"
	"os"
)

func sol1() {
	var size = 14
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(f)-size-1; i++ {
		if checkUniq(f[i : i+size]) {
			fmt.Println(i + size)
			break
		}
	}
}

func checkUniq(substring []byte) bool {
	fmt.Printf("%s\n", substring)
	m := make(map[byte]interface{})
	for i := range substring {
		_, ok := m[substring[i]]
		if ok {
			return false
		}
		m[substring[i]] = nil
	}
	return true
}
