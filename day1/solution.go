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

// add all the input together
func solvePartOne(lines []string) {
	total := 0

	for _, value := range lines {
		number, convertErr := strconv.Atoi(value)
		if convertErr == nil {
			total += number
		}
	}
    
    fmt.Print("PART 1) " + strconv.Itoa(total) + "\n")
}

// add the input into a running total until it reaches the same running total twice
func solvePartTwo(lines []string) {
	runningTotal := 0
	notFound := true
	foundValues :=make(map[int]bool) 

	for notFound {
		for _, value := range lines {
			number, convertErr := strconv.Atoi(value)

			if convertErr == nil {
				runningTotal += number

				if foundValues[runningTotal] {
					fmt.Print("PART 2) " + strconv.Itoa(runningTotal) + "\n")
					notFound = false
					break
				}

				foundValues[runningTotal] = true
			}
		}
	}
}