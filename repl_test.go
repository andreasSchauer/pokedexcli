package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input 		string
		expected 	[]string
	}{
		{
			input:		"  hello  world  ",
			expected:	[]string{"hello", "world"},
		},
		{
			input: 		"hello ",
			expected: 	[]string{"hello"},
		},
		{
			input: 		"HELLO ",
			expected: 	[]string{"hello"},
		},
		{
			input: 		"",
			expected: 	[]string{},
		},
		{
			input: 		"     ",
			expected: 	[]string{},
		},
		{
			input: 		"Charmander BuLbASaUr   PIKACHU ",
			expected: 	[]string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	
	for i, tc := range tests {
		actual := cleanInput(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Testcase %d: words don't match for input: %s. expected: %v, actual: %v", i, tc.input, tc.expected, actual)
		}
	}

}