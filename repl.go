package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var split []string
	words := strings.Fields(text)
	for i := range words {
		split = append(split, strings.ToLower(words[i]))
	}
	return split
}