package main

import (
	"regexp"
	"strings"
)

// Tokenize splits a string into tokens
func Tokenize(s string) []string {
	re := regexp.MustCompile("\\s+")
	return re.Split(s, -1)
}

// TokenizeLowerCase splits a string into lowercase tokens
func TokenizeLowerCase(s string) []string {
	tokens := Tokenize(s)
	for i := range tokens {
		tokens[i] = strings.ToLower(tokens[i])
	}
	return tokens
}
