package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) push(i ...rune) {
	s.items = append(s.items, i...)
}

func (s *Stack) pop(amount int) []rune {
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

func (s *Stack) peep() rune {
	length := len(s.items)
	if length == 0 {
		return ' '
	}
	return s.items[length-1]
}

func (s *Stack) printStack() {
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

func main() {
	// sol1()
	sol2()
}
