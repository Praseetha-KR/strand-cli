package cmd

import (
	"log"
)

func validatedLength(l int) int {
	if l <= 0 {
		log.Fatal("Invalid value for length flag: must be greater than 0")
	}
	return l
}
