package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (s *Stack) pop2(amount int) []rune {
	length := len(s.items)
	if length == 0 {
		return []rune{}
	}
	final := s.items[length-amount:]
	// if amount != 1 {
	// 	final = reverseStack(final)
	// }
	s.items = s.items[:length-amount]
	return final
}

func sol2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buffer := make([]byte, 4)
	stackNum := 0
	stacks := make([]Stack, 9)
	for {
		_, err := f.Read(buffer)
		if err != nil {
			log.Print(err)
			break
		}
		if rune(buffer[0]) == '\n' {
			break // reached end of Stack definition
		}
		char := rune(buffer[1])
		if char >= 65 && char <= 90 {
			stacks[stackNum].items = append(stacks[stackNum].items, char)
		}
		stackNum += 1
		if rune(buffer[3]) == '\n' {
			stackNum = 0 // if the last char is a new line, start filling from 1st Stack again
		}

	}
	for i := range stacks {
		stacks[i].items = reverseStack(stacks[i].items)
	}

	// fix the first three bytes that the previous step consumed
	var invalidbuffer []byte
	invalidbuffer = []byte("mov")
	remaining, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	invalidbuffer = append(invalidbuffer, remaining...)
	//loop over each line and extract quantities and locations
	sc := bufio.NewScanner(bytes.NewReader(invalidbuffer))
	for sc.Scan() {
		var to, from, amount int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amount, &from, &to)
		stacks[to-1].push(stacks[from-1].pop2(amount)...)
	}
	for _, i := range stacks {
		fmt.Print(string(i.peep()))
	}
}
