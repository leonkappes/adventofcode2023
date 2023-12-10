package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var DigitMap = map[string]string{
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

func main() {
	b, err := os.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(b), "\n")

	var lineDigitsPart1 []int
	var lineDigitsPart2 []int

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		lineDigitsPart1 = append(lineDigitsPart1, RetrieveDigits(currentLine))
		currentLine = ReplaceTextWithDigits(currentLine)
		lineDigitsPart2 = append(lineDigitsPart2, RetrieveDigits(currentLine))
	}

	puzzleTotalPart1 := 0
	puzzleTotalPart2 := 0
	for i := 0; i < len(lineDigitsPart1); i++ {
		puzzleTotalPart1 += lineDigitsPart1[i]
		puzzleTotalPart2 += lineDigitsPart2[i]
	}

	fmt.Printf("[Part1] Total Sum is %d\n", puzzleTotalPart1)
	fmt.Printf("[Part2] Total Sum is %d\n", puzzleTotalPart2)
}

func RetrieveDigits(line string) int {
	allNumbers := regexp.MustCompile("\\d").FindAllString(line, -1)
	thisLineValue, _ := strconv.Atoi(allNumbers[0] + allNumbers[len(allNumbers)-1])
	return thisLineValue
}

func ReplaceTextWithDigits(line string) string {
	for k, v := range DigitMap {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}
