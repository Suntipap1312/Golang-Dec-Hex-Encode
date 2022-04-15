package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if arges := os.Args[1:]; len(arges) == 2 {
		start, _ = strconv.Atoi(arges[0])
		stop, _ = strconv.Atoi(arges[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec", "hex", "encode", strings.Repeat("-", 45))

	for n := start; n < stop; n++ { // %c => print a character depending on the given code point
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}
