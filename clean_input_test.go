package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name		string
		input 		string
		expected 	[]string
	}{
		{
			name:		"lots of whitespace",
			input:		"  hello  world  ",
			expected:	[]string{"hello", "world"},
		},
		{
			name:		"only one word",
			input: 		"hello ",
			expected: 	[]string{"hello"},
		},
		{
			name:		"all caps",
			input: 		"HELLO ",
			expected: 	[]string{"hello"},
		},
		{
			name:		"empty string",
			input: 		"",
			expected: 	[]string{},
		},
		{
			name:		"only whitespace",
			input: 		"     ",
			expected: 	[]string{},
		},
		{
			name:		"different caps",
			input: 		"Charmander BuLbASaUr   PIKACHU ",
			expected: 	[]string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	
	for i, tc := range tests {
		actual := cleanInput(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Error on testcase %d - %s: words don't match for input: %s. expected: %v, actual: %v", i, tc.name, tc.input, tc.expected, actual)
		}
	}

}