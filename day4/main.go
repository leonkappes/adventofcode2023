package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

var cardMap map[int]*Card

type Card struct {
	winningNumbers  []string
	numbers         []string
	totalPoints     int
	id              int
	instances       int
	matchingNumbers int
}

func (c *Card) CountPoints() int {
	numbers := 0
	for _, number := range c.numbers {
		if slices.Contains(c.winningNumbers, number) {
			numbers += 1
		}
	}

	c.matchingNumbers = numbers

	if numbers > 1 {
		return int(math.Pow(2, float64(numbers-1)))
	} else {
		return numbers
	}
}

func (c *Card) CalculateWonCards() {
	for i := 0; i < c.matchingNumbers; i++ {
		if cardMap[c.id+i+1] != nil {
			cardMap[c.id+i+1].instances += c.instances
		}
	}
}

func main() {
	b, err := os.ReadFile("day4/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(b), "\n")

	cardMap = make(map[int]*Card)

	totalPoints := 0

	for id, line := range lines {
		winningNumbers := slices.DeleteFunc(strings.Split(line[strings.IndexByte(line, ':')+1:strings.IndexByte(line, '|')], " "), func(e string) bool {
			return e == "" || e == " "
		})
		ownNumbers := slices.DeleteFunc(strings.Split(line[strings.IndexByte(line, '|')+1:], " "), func(e string) bool {
			return e == "" || e == " "
		})
		card := &Card{
			id:              id + 1,
			numbers:         ownNumbers,
			winningNumbers:  winningNumbers,
			totalPoints:     0,
			instances:       1,
			matchingNumbers: 0,
		}
		card.totalPoints = card.CountPoints()
		totalPoints += card.totalPoints
		cardMap[card.id] = card
	}

	fmt.Printf("[Part 1] Total Points: %d\n", totalPoints)

	for i := 1; i < len(cardMap); i++ {
		cardMap[i].CalculateWonCards()
	}

	cardSum := 0
	for _, card := range cardMap {
		cardSum += card.instances
	}

	fmt.Printf("[Part 2] Total Cards: %d\n", cardSum)
}
