package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Direction string

const (
	left  Direction = "left"
	right Direction = "right"
	up    Direction = "up"
	down  Direction = "down"
)

var tileGrid [][]*Tile

type Tile struct {
	x             int
	y             int
	char          string
	key           string
	outDirections []Direction
	inDirections  []Direction
}

func NewTile(x, y int, char string) *Tile {
	tile := &Tile{
		x:    x,
		y:    y,
		char: char,
		key:  strconv.Itoa(x) + strconv.Itoa(y),
	}
	tile.calculateDirections()
	return tile
}

func (t *Tile) calculateDirections() {
	switch t.char {
	case "|":
		t.outDirections = []Direction{up, down}
		t.inDirections = []Direction{down, up}
	case "-":
		t.outDirections = []Direction{left, right}
		t.inDirections = []Direction{right, left}
	case "L":
		t.outDirections = []Direction{up, right}
		t.inDirections = []Direction{down, left}
	case "J":
		t.outDirections = []Direction{up, left}
		t.inDirections = []Direction{down, right}
	case "7":
		t.outDirections = []Direction{down, left}
		t.inDirections = []Direction{up, right}
	case "F":
		t.outDirections = []Direction{down, right}
		t.inDirections = []Direction{up, left}
	case ".":
		t.outDirections = []Direction{}
		t.inDirections = []Direction{}
	case "S":
		t.outDirections = []Direction{up, down, left, right}
		t.inDirections = []Direction{up, down, left, right}

	}
}

func isInBounds(x, y int) bool {
	return x < len(tileGrid) && y < len(tileGrid[x]) && x >= 0 && y >= 0
}

func findNeighbours(tile *Tile) []*Tile {
	var neighbours []*Tile

	if tile.x == 34 && tile.y == 138 {
		println("Here")
	}

	for _, direction := range tile.outDirections {
		switch direction {
		case down:
			if isInBounds(tile.x-1, tile.y) {
				tileT := tileGrid[tile.x-1][tile.y]
				if slices.Contains(tileT.inDirections, down) {
					neighbours = append(neighbours, tileT)
				}
			}
		case up:
			if isInBounds(tile.x+1, tile.y) {
				tileT := tileGrid[tile.x+1][tile.y]
				if slices.Contains(tileT.inDirections, up) {
					neighbours = append(neighbours, tileT)
				}
			}
		case left:
			if isInBounds(tile.x, tile.y-1) {
				tileT := tileGrid[tile.x][tile.y-1]
				if slices.Contains(tileT.inDirections, left) {
					neighbours = append(neighbours, tileT)
				}
			}
		case right:
			if isInBounds(tile.x, tile.y+1) {
				tileT := tileGrid[tile.x][tile.y+1]
				if slices.Contains(tileT.inDirections, right) {
					neighbours = append(neighbours, tileT)
				}
			}
		}
	}

	return neighbours

}

func main() {
	b, err := os.ReadFile("day10/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	var startTile *Tile
	lines := strings.Split(string(b), "\n")
	slices.Reverse(lines)
	for row := 0; row < len(lines); row++ {
		var currentRow []*Tile
		for column := 0; column < len(lines[row]); column++ {
			tile := NewTile(row, column, lines[row][column:column+1])
			if tile.char == "S" {
				startTile = tile
			}
			currentRow = append(currentRow, tile)
		}
		tileGrid = append(tileGrid, currentRow)
	}

	// Do BFS
	seen := make([][]bool, len(tileGrid))
	for i := range seen {
		seen[i] = make([]bool, len(tileGrid[i]))
	}
	loop := []*Tile{startTile}
	seen[startTile.x][startTile.y] = true

	for {
		top := loop[len(loop)-1]
		neighbours := findNeighbours(top)
		if len(neighbours) != 2 {
			fmt.Printf("Found position with %d neigbours: %v %v\n", len(neighbours), top, neighbours[0])
			panic("")
		}

		// Pop every seen neighbour
		for len(neighbours) > 0 && seen[neighbours[0].x][neighbours[0].y] {
			neighbours = neighbours[1:]
		}

		if len(neighbours) == 0 {
			break
		}

		loop = append(loop, neighbours[0])
		seen[neighbours[0].x][neighbours[0].y] = true
	}

	fmt.Printf("[Part1] Furthest away point is %d Steps away\n", len(loop)/2)

	// A = Area I = Points in Polygon(what we want) R = Points on the Polygon edges
	// A = I + R/2 - 1
	// I = A + (-R)/2 + 1
	sum := 0
	for i := range loop {
		n := i + 1
		if n == len(loop) {
			n = 0
		}
		sum += loop[i].x*loop[n].y - loop[n].x*loop[i].y
	}

	a := int(math.Abs(float64(sum)) / 2)

	pointsInPolygon := a + (len(loop)*-1)/2 + 1

	fmt.Printf("[Part2] Points in Polygon is %d \n", pointsInPolygon)

}
