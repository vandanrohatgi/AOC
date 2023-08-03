package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type stack struct {
	items []rune
}

func (s *stack) push(i ...rune) {
	s.items = append(s.items, i...)
}

func (s *stack) pop(amount int) []rune {
	length := len(s.items)
	if length == 0 {
		return []rune{}
	}
	final := s.items[length-amount:]
	if amount != 1 {
		final = reverseStack(final)
	}
	s.items = s.items[:length-amount]
	return final
}

func (s *stack) peep() rune {
	length := len(s.items)
	if length == 0 {
		return ' '
	}
	return s.items[length-1]
}

func (s *stack) printStack() {
	for _, i := range s.items {
		fmt.Println(string(i))
	}
}

func reverseStack(final []rune) []rune {
	for i, j := 0, len(final)-1; i < j; i, j = i+1, j-1 {
		final[i], final[j] = final[j], final[i]
	}
	return final

}

func sol1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buffer := make([]byte, 4)
	stackNum := 0
	stacks := make([]stack, 9)
	for {
		_, err := f.Read(buffer)
		if err != nil {
			log.Print(err)
			break
		}
		if rune(buffer[0]) == '\n' {
			break // reached end of stack definition
		}
		char := rune(buffer[1])
		if char >= 65 && char <= 90 {
			stacks[stackNum].items = append(stacks[stackNum].items, char)
		}
		stackNum += 1
		if rune(buffer[3]) == '\n' {
			stackNum = 0 // if the last char is a new line, start filling from 1st stack again
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
		// fmt.Println(sc.Text())
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amount, &from, &to)
		// fmt.Println(to, from, amount)
		stacks[to-1].push(stacks[from-1].pop(amount)...)
	}
	for _, i := range stacks {
		fmt.Print(string(i.peep()))
	}
}
