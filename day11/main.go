package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var lines []string

type Galaxy struct {
	x int
	y int
}

func main() {
	b, err := os.ReadFile("day11/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines = strings.Split(string(b), "\n")

	var galaxyArray []Galaxy
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			curChar := string(lines[i][j])
			if curChar == "#" {
				galaxyArray = append(galaxyArray, Galaxy{x: i, y: j})
			}
		}
	}

	newGalaxies := FillEmptyLines(galaxyArray, 1)
	sum := 0
	for i := range newGalaxies {
		for j := i + 1; j < len(newGalaxies); j++ {
			dist := getDistance(newGalaxies[i], newGalaxies[j])
			sum += dist
		}
	}

	fmt.Printf("[Part1] Sum of Distances is %d\n", sum)

	newGalaxies2 := FillEmptyLines(galaxyArray, 999_999)
	sum = 0
	for i := range newGalaxies2 {
		for j := i + 1; j < len(newGalaxies2); j++ {
			dist := getDistance(newGalaxies2[i], newGalaxies2[j])
			sum += dist
		}
	}

	fmt.Printf("[Part2] Sum of Distances is %d\n", sum)
}

func FillEmptyLines(galaxies []Galaxy, factor int) []Galaxy {
	galaxyRows := make(map[int]bool)
	galaxyColumns := make(map[int]bool)
	for _, galaxy := range galaxies {
		galaxyRows[galaxy.x] = true
		galaxyColumns[galaxy.y] = true
	}

	var rowsWithoutGalaxies []int
	for row := 0; row <= len(lines); row++ {
		if !galaxyRows[row] {
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, row)
		}
	}
	var colsWithoutGalaxies []int
	for col := 0; col <= len(lines[0]); col++ {
		if !galaxyColumns[col] {
			colsWithoutGalaxies = append(colsWithoutGalaxies, col)
		}
	}

	var newGalaxies []Galaxy

	for _, galaxy := range galaxies {
		c := galaxy
		for _, row := range rowsWithoutGalaxies {
			if row > galaxy.x {
				break
			}
			c.x += factor
		}
		for _, col := range colsWithoutGalaxies {
			if col > galaxy.y {
				break
			}
			c.y += factor
		}
		newGalaxies = append(newGalaxies, c)
	}

	return newGalaxies
}

func getDistance(a, b Galaxy) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}
