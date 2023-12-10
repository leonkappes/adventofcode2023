package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day9/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")

	puzzleTotal := 0
	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		values := strings.Fields(currentLine)
		lastValueInt, _ := strconv.Atoi(values[len(values)-1])
		lastDifference := lastValueInt

		done := false
		var differences []int
		for j := 1; j < len(values); j++ {
			lastValue, _ := strconv.Atoi(values[j-1])
			currentValue, _ := strconv.Atoi(values[j])
			difference := currentValue - lastValue
			differences = append(differences, difference)
		}
		lastDifference += differences[len(differences)-1]
		for !done {
			var currentDifference []int
			for j := 1; j < len(differences); j++ {
				lastValue := differences[j-1]
				currentValue := differences[j]
				difference := currentValue - lastValue
				currentDifference = append(currentDifference, difference)
			}
			lastDifference += currentDifference[len(currentDifference)-1]
			differences = nil
			differences = append(differences, currentDifference...)
			nullCounter := 0
			for l := 0; l < len(currentDifference); l++ {
				cur := currentDifference[l]
				if cur == 0 {
					nullCounter++
				}
			}
			if nullCounter == len(currentDifference) {
				done = !done
			}
		}
		puzzleTotal += lastDifference
	}

	fmt.Printf("[Part1] Solution: %d\n", puzzleTotal)

	puzzleTotal = 0

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		values := strings.Fields(currentLine)
		lastValueInt, _ := strconv.Atoi(values[0])
		var firstRow []int
		firstRow = append(firstRow, lastValueInt)

		done := false
		var differences []int
		for j := 1; j < len(values); j++ {
			lastValue, _ := strconv.Atoi(values[j-1])
			currentValue, _ := strconv.Atoi(values[j])
			difference := currentValue - lastValue
			differences = append(differences, difference)
		}
		firstRow = append(firstRow, differences[0])
		for !done {
			var currentDifference []int
			for j := 1; j < len(differences); j++ {
				lastValue := differences[j-1]
				currentValue := differences[j]
				difference := currentValue - lastValue
				currentDifference = append(currentDifference, difference)
			}
			firstRow = append(firstRow, currentDifference[0])
			differences = nil
			differences = append(differences, currentDifference...)
			nullCounter := 0
			for l := 0; l < len(currentDifference); l++ {
				cur := currentDifference[l]
				if cur == 0 {
					nullCounter++
				}
			}
			if nullCounter == len(currentDifference) {
				done = !done
			}
		}

		slices.Reverse(firstRow)
		prediction := 0
		for k := 0; k < len(firstRow); k++ {
			prediction = firstRow[k] - prediction
		}
		puzzleTotal += prediction
	}

	fmt.Printf("[Part2] Solution: %d\n", puzzleTotal)

}
