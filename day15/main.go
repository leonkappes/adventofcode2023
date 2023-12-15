package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var boxes [][]BoxEntry

func main() {
	b, err := os.ReadFile("day15/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	rounds := strings.Split(string(b), ",")
	sum := int32(0)
	for _, round := range rounds {
		currentValue := int32(0)
		for _, char := range round {
			currentValue += char
			currentValue *= 17
			currentValue = currentValue % 256
		}
		sum += currentValue
	}

	fmt.Printf("[Part1] Sum is %d\n", sum)

	boxes = make([][]BoxEntry, 256)
	for _, round := range rounds {
		splitRound := regexp.MustCompile("[-=]").Split(round, -1)
		name := splitRound[0]
		boxIndex := int32(0)
		for _, char := range name {
			boxIndex += char
			boxIndex *= 17
			boxIndex = boxIndex % 256
		}
		focal := splitRound[1]
		if strings.Contains(round, "=") {
			ReplaceOrAppendEntry(boxIndex, name, focal)
		} else if strings.Contains(round, "-") {
			RemoveEntry(boxIndex, name)
		}
	}

	sum2 := 0
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i]); j++ {
			sum2 += (i + 1) * (j + 1) * boxes[i][j].focalLength
		}
	}

	fmt.Printf("[Part2] Sum is %d\n", sum2)

}

func RemoveEntry(index int32, name string) {
	box := boxes[index]
	entryToRemove := -1
	for i, entry := range box {
		if entry.name == name {
			entryToRemove = i
		}
	}
	if entryToRemove != -1 {
		boxes[index] = append(boxes[index][:entryToRemove], boxes[index][entryToRemove+1:]...)
	}

}

func ReplaceOrAppendEntry(index int32, name string, focal string) {
	box := boxes[index]
	focalInt, _ := strconv.Atoi(focal)
	success := false
	for i, entry := range box {
		if entry.name == name {
			boxes[index][i].focalLength = focalInt
			success = true
		}
	}

	if !success {
		boxes[index] = append(boxes[index], BoxEntry{
			name:        name,
			focalLength: focalInt,
		})
	}
}

type BoxEntry struct {
	name        string
	focalLength int
}
