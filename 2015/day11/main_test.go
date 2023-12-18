package main

import "testing"

type Tests struct {
	Name  string
	Input string
	Want  bool
}

func TestHasTwoPairs(t *testing.T) {
	tests := []Tests{
		{"valid", "aabb", true},
		{"invalid", "aabgfaa", false},
		{"invalid", "aabcdf", false},
	}

	for _, i := range tests {
		output := hasTwoPairs(i.Input)
		if output != i.Want {
			t.Fatalf("%s expected: %t, received :%t", i.Name, i.Want, output)
		}
	}

}

func TestNotHaveBadLetters(t *testing.T) {
	tests := []Tests{
		{"valid", "abcd", true},
		{"invalid", "aeiou", false},
	}
	for _, i := range tests {
		output := notHaveBadLetters(i.Input)
		if output != i.Want {
			t.Fatalf("%s expected: %t, received :%t", i.Name, i.Want, output)
		}
	}
}

func TestHasThreeLetters(t *testing.T) {
	tests := []Tests{
		{"valid", "abcd", true},
		{"invalid", "aeiou", false},
	}
	for _, i := range tests {
		output := hasThreeLetters(i.Input)
		if output != i.Want {
			t.Fatalf("%s expected: %t, received :%t", i.Name, i.Want, output)
		}
	}
}

func Test_helper(t *testing.T) {
	var s santaPass

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"increment by one", "abcd", "abce"},
		{"wrap around", "abcz", "abda"},
		{"ripple change", "azzz", "baaa"},
		{"no change", "zzzz", "zzzz"}, // edge case which im not handling
	}

	for _, i := range tests {
		s.setPass(i.input)
		s.incrementByOne(len(i.input) - 1)
		if s.getString() != i.want {
			t.Fatalf(" %s expected %s, received %s", i.name, i.want, s.getString())
		}
	}
}
