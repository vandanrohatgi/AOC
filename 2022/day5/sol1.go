package main

import (
	"fmt"
	"log"
	"os"
)

type stack struct {
	items []rune
}

func (s *stack) push(i rune) {
	s.items = append(s.items, i)
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

func reverseStack(final []rune) []rune {
	for i, j := 0, len(final)-1; i < j; i, j = i+1, j-1 {
		final[i], final[j] = final[j], final[i]
	}
	return final

}

func sol1() {
	f, err := os.Open("test.txt")
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
			log.Fatal(err)
		}
		if string(buffer[0]) == "\n" {
			break
		}
		if string(buffer[3]) == "\n" {
			stackNum = 0
		}
		char := rune(buffer[1])
		if char <= 90 && char >= 65 {
			stacks[stackNum].items = append(stacks[stackNum].items, char)
		}
		stackNum += 1
	}
	// sc := bufio.NewScanner(f)
	// sc := bufio.NewReader(f)
	// sc.Split(bufio.ScanLines)
	// for sc.Scan() && sc.Text() != "" {
	// fmt.Println(sc.
	// }
	//    var previousRune rune
	//    var spaceCount int
	// for {
	// 	if c, _, err := sc.ReadRune(); err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			log.Fatal(err)
	// 		}
	// 	} else {
	//            if previousRune==c && c=='\n'{
	//                break
	//            }
	//            previousRune=c
	//            if c==' '{
	//                spaceCount+=1
	//            }
	//            if c>=65 && c<=90{
	//                if spaceCount == 1
	//            }
	// 		fmt.Printf("%s", string(c))
	// 	}
	// }
	fmt.Print(stacks)
}
