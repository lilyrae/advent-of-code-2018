package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	fileName := "input.txt"
	lines := readFile(fileName)
	solvePartOne(lines)
	solvePartTwo(lines)	
}

func readFile(fileName string) (lines []string) {
	input, readErr := ioutil.ReadFile(fileName)
	if readErr != nil {
        panic(readErr)
    }

	lines = strings.Split(string(input), "\n")
	return
}

// find how many times of a letter occurs exactly twice and exactly three time
func solvePartOne(lines []string) {
	containsTwo := 0
	containsThree := 0

	for _, line := range lines {
		letterCount :=make(map[string]int)

		for _, char := range line {
			letterCount[string(char)]++
		}

		if hasLetterCount(2, letterCount) {
			containsTwo++
		}

		if hasLetterCount(3, letterCount) {
			containsThree++
		}
	}
    
    fmt.Print("PART 1) " + strconv.Itoa(containsTwo * containsThree) + "\n")
}

func hasLetterCount(desiredCount int, letterCount map[string]int) (hasCount bool) {
	hasCount = false
	
	for _, count := range letterCount {
		if count == desiredCount {
			hasCount = true
			return
		}
	}
	return
}

// find two strings with the difference of one letter
func solvePartTwo(lines []string) {
	lineA, lineB := findLinesWithOneDifferece(lines)
	sharedSeq := findSharedCharacters(lineA, lineB)

	fmt.Print("PART 2) " + sharedSeq + "\n")
}

func findLinesWithOneDifferece(lines []string) (lineA string, lineB string) {
	for _, line := range lines {
		lineA = line

		for _, lineToCompare := range lines {
			lineB = lineToCompare
			differences := getLineDifferences(lineA, lineB)

			if differences == 1 {
				return
			}
		}
	}
	return
}

func getLineDifferences(a string, b string) (differences int) {
	differences = 0
	
	if len(a) != len(b) {
		differences = -1
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}
	return
}

func findSharedCharacters(a string, b string) (sharedSequence string) {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			sharedSequence += string(a[i])
		}
	}
	return
}