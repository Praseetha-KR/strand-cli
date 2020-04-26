package main

import (
	"log"
	"strings"
)

func validatedLength(l int) int {
	if l <= 0 {
		log.Fatal("Invalid value for length flag: must be greater than 0")
	}
	return l
}

func validatedCase(c string) int {
	switch alphacase := strings.ToLower(c); alphacase {
	case
		"lower",
		"lowercase",
		"l":
		return LOWER
	case
		"upper",
		"uppercase",
		"u":
		return UPPER
	case
		"all",
		"a",
		"":
		return ALL
	default:
		log.Fatal("Invalid value for case flag: supported are lower|upper|L|U")
	}
	return ALL
}
