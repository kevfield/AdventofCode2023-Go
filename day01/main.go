package main

import (
	"fmt"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	fmt.Println(puzzleInput)
}
