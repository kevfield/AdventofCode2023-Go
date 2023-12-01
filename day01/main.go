package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	fmt.Println("Part 1 = ", getfirstLast(puzzleInput))
}

func getfirstLast(getfLinput []string) int {
	combinestrings := ""
	incCount := 0
	totalCount := 0
	var firstInt, lastInt string
	for i := 0; i < len(getfLinput); i++ {
		// filter out ints in the imported string
		regFilter := regexp.MustCompile("[0-9]")
		currentInt := (regFilter.FindAllString(getfLinput[i], -1))
		firstInt = currentInt[0]
		lastInt = currentInt[len(currentInt)-1]
		combinestrings = fmt.Sprintf("%s%s", firstInt, lastInt)
		incCount, _ = strconv.Atoi(combinestrings)
		totalCount = totalCount + incCount
	}
	return totalCount

}
