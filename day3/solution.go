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
	fabric := findFabricPositions(lines)
	solvePartOne(fabric)
	solvePartTwo(lines, fabric)	
}

func readFile(fileName string) (lines []string) {
	input, readErr := ioutil.ReadFile(fileName)
	if readErr != nil {
        panic(readErr)
    }

	lines = strings.Split(string(input), "\n")
	return
}

// find how many times fabric overlaps (positions repeat)
func solvePartOne(fabric map[string]int) {
	totalOverlaps := 0

	for _, count := range fabric {
		if count > 1 {
			totalOverlaps++
		}
	}
    
    fmt.Print("PART 1) " + strconv.Itoa(totalOverlaps) + "\n")
}

// find fabric portion with no overlaps
func solvePartTwo(lines []string, fabric map[string]int) {
	fabricId := findFabricWithoutOverlaps(lines, fabric)
    fmt.Print("PART 2) " + fabricId + "\n")
}

func findFabricPositions(lines []string) (fabric map[string]int) {
	fabric = make(map[string]int)

	for _, line := range lines {
		_, x, y, width, height := parseInput(line)

		for indexX := 0; indexX < width; indexX++ {
			for indexY := 0; indexY < height; indexY++ {
				key := buildKey(x, y, indexX, indexY)
				fabric[key]++
			}
		}
	}
	return
}

func findFabricWithoutOverlaps(lines []string, fabric map[string]int) (lineID string) {
	for _, line := range lines {
		id, x, y, width, height := parseInput(line)
		hasOverlap := false

		for indexX := 0; indexX < width; indexX++ {
			for indexY := 0; indexY < height; indexY++ {
				key := buildKey(x, y, indexX, indexY)
				
				if val, ok := fabric[key]; ok {
					if val > 1 {
						hasOverlap = true
					}
				}
			}
		}

		if !hasOverlap {
			lineID = id
			return
		}
	}
	return
}

func parseInput(line string) (lineID string, x int, y int, width int, height int) {
	values := strings.Split(line, " ")
	lineID = values[0]
	startingPoints := strings.Split(values[2], ",")

	x, err := strconv.Atoi(startingPoints[0])
	if err != nil { log.Fatal(err) }

	y, err = strconv.Atoi(strings.Split(startingPoints[1], ":")[0])
	if err != nil { log.Fatal(err) }

	lengths := strings.Split(values[3], "x")
	width, err = strconv.Atoi(lengths[0])
	if err != nil { log.Fatal(err) }

	height, err = strconv.Atoi(lengths[1])
	if err != nil { log.Fatal(err) }

	return
}

func buildKey(x int, y int, indexX int, indexY int) (key string) {
	key = strconv.Itoa(x + indexX) + "-" + strconv.Itoa(y + indexY)
	return
}