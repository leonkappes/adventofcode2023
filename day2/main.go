package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id      int
	isValid bool
	gameMap []map[string]int
	power   int
}

func NewGame(line string) *Game {
	id, _ := strconv.Atoi(line[strings.IndexByte(line, 'e')+2 : strings.IndexByte(line, ':')])
	rounds := strings.Split(line[strings.IndexByte(line, ':')+2:], "; ")

	var gameMapList []map[string]int
	for _, round := range rounds {
		colorSplit := strings.Split(round, ", ")
		roundMap := make(map[string]int)
		for _, colorValueCombo := range colorSplit {
			separated := strings.Fields(colorValueCombo)
			color := separated[1]
			value, _ := strconv.Atoi(separated[0])
			roundMap[color] = value
		}
		gameMapList = append(gameMapList, roundMap)
	}
	game := &Game{
		id:      id,
		gameMap: gameMapList,
	}
	game.calculateValid()
	game.calculatePower()
	return game
}

func (g *Game) calculateValid() {
	stillValid := true
	for i := 0; i < len(g.gameMap) && stillValid; i++ {
		round := g.gameMap[i]
		for color, value := range round {
			switch color {
			case "red":
				if value > 12 {
					stillValid = false
					break
				}
			case "green":
				if value > 13 {
					stillValid = false
					break
				}
			case "blue":
				if value > 14 {
					stillValid = false
					break
				}
			}
		}
	}

	g.isValid = stillValid
}

func (g *Game) calculatePower() {
	fewest := make(map[string]int)
	for i := 0; i < len(g.gameMap); i++ {
		round := g.gameMap[i]
		for color, value := range round {
			current := 0
			if _, ok := fewest[color]; ok {
				current = fewest[color]
			}
			fewest[color] = max(value, current)
		}
	}

	product := 0
	for _, i := range fewest {
		if product == 0 {
			product = i
		} else {
			product *= i
		}
	}

	g.power = product
}

func main() {
	b, err := os.ReadFile("day2/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")
	var games []*Game
	gameSum := 0
	gamePowerSum := 0

	for i := 0; i < len(lines); i++ {
		game := NewGame(lines[i])
		games = append(games, game)
		if game.isValid {
			gameSum += game.id
		}
		gamePowerSum += game.power
	}

	fmt.Printf("[Part1] Total GameSum is %d\n", gameSum)
	fmt.Printf("[Part2] Total GameSum is %d\n", gamePowerSum)
}
