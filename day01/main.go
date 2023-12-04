package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	fmt.Println("Part 1 = ", getfirstLast(puzzleInput))
	part2Calc := detectintsasText(puzzleInput)
	fmt.Println("Part2 = ", getfirstLast(part2Calc))

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

func detectintsasText(detectInput []string) []string {
	var newSlice []string
	var newInput, newOutput string
	foundSomething := 0
	numberMapping := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for i := 0; i < len(detectInput); i++ {
		newInput = detectInput[i]
		for key, val := range numberMapping {
			if strings.Contains(newInput, key) {
				//newOutput = strings.Replace(newInput, key, val, -1)
				newOutput = strings.ReplaceAll(newInput, key, val)
				newInput = newOutput
				foundSomething++
			}
		}
		if foundSomething == 0 {
			newSlice = append(newSlice, newInput)
		} else {
			foundSomething = 0
			newSlice = append(newSlice, newOutput)
		}
	}

	return newSlice
}
