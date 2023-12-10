package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day6/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")
	times := strings.FieldsFunc(lines[0][strings.IndexByte(lines[0], ':')+1:], func(r rune) bool {
		return r == ' '
	})
	distance := strings.FieldsFunc(lines[1][strings.IndexByte(lines[1], ':')+1:], func(r rune) bool {
		return r == ' '
	})

	puzzleValue := 1
	for i := 0; i < len(times); i++ {
		race := ParseRace(times[i], distance[i])
		puzzleValue *= race.validReleasePoints
	}

	fmt.Printf("[Part1] Solution is %d\n", puzzleValue)

	timesPart2 := ""
	for i := 0; i < len(times); i++ {
		timesPart2 += times[i]
	}
	distancePart2 := ""
	for i := 0; i < len(distance); i++ {
		distancePart2 += distance[i]
	}
	puzzleValue = 0
	race := ParseRace(timesPart2, distancePart2)
	puzzleValue = race.validReleasePoints

	fmt.Printf("[Part2] Solution is %d\n", puzzleValue)
}

type Race struct {
	time               int
	distance           int
	validReleasePoints int
}

func ParseRace(time string, distance string) Race {
	timeInt, _ := strconv.Atoi(time)
	distanceInt, _ := strconv.Atoi(distance)

	validReleasePoints := 0
	for i := 0; i < timeInt; i++ {
		y := CalculateFunctionValue(i, timeInt)
		if y > distanceInt {
			validReleasePoints += 1
		}
	}

	return Race{
		time:               timeInt,
		distance:           distanceInt,
		validReleasePoints: validReleasePoints,
	}
}

func CalculateFunctionValue(x int, time int) int {
	// -x^2+(time)x
	bx := time * x
	ax := math.Pow(float64(x), 2) * -1

	return int(ax) + bx
}
