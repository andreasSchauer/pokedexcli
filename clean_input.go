package main

import (
	"strings"

)

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	words := strings.Fields(textLower)

	return words
}