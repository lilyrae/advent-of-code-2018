package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
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

	lines = strings.Split(string(input), " ")
	return
}

func solvePartOne(lines []string) {
	_, result := getMeta(lines)

	fmt.Print("PART 1) " + strconv.Itoa(result) + "\n")
}

func getMeta(input []string) (remainingInput []string, metaTotal int) {
	children, metaQuantity, remainingInput := parseInput(input)
	metaTotal = 0
	childTotal := 0

	for children > 0 {
		remainingInput, childTotal = getMeta(remainingInput)
		metaTotal += childTotal
		children--
	}

	metaTotal += sumMetaValues(metaQuantity, remainingInput)
	remainingInput = remainingInput[metaQuantity:]
	return
}


func solvePartTwo(lines []string) {
	_, result := getChildMeta(lines)

	fmt.Print("PART 2) " + strconv.Itoa(result) + "\n")
}

func getChildMeta(input []string) (remainingInput []string, metaTotal int) {
	children, metaQuantity, remainingInput := parseInput(input)
	childMeta := make(map[int]int)
	childTotal := 0
	childCount := 0
	metaTotal = 0

	// add together the meta values of child nodes
	for childCount < children {
		childCount++
		remainingInput, childTotal = getChildMeta(remainingInput)
		childMeta[childCount] = childTotal
	}

	if children > 0 {
		metaTotal += sumReferenceChildMetaValues(metaQuantity, remainingInput, childMeta)
	} else {
		metaTotal += sumMetaValues(metaQuantity, remainingInput)	
	}

	remainingInput = remainingInput[metaQuantity:]
	return
}

func parseInput(input []string) (children int, metaQuantity int, remainingInput []string) {
	children, err := strconv.Atoi(input[0])
	if err != nil { log.Fatal(err) }

	metaQuantity, err = strconv.Atoi(input[1])
	if err != nil { log.Fatal(err) }

	remainingInput = input[2:]
	return
}

func sumMetaValues(metaQuantity int, remainingInput []string) (metaTotal int) {
	metaTotal = 0
	count := 0

	for count < metaQuantity {
		newMeta, err := strconv.Atoi(remainingInput[count])
		if err != nil { log.Fatal(err) }

		metaTotal += newMeta
		count++
	}
	return
}

func sumReferenceChildMetaValues(metaQuantity int, remainingInput []string, childMeta map[int]int) (metaTotal int) {
	metaTotal = 0
	count := 0

	for count < metaQuantity {
		newMetaIndex, err := strconv.Atoi(remainingInput[count])
		if err != nil { log.Fatal(err) }

		if val, ok := childMeta[newMetaIndex]; ok {
			metaTotal += val
		}
		count++
	}
	return
}