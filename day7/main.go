package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	kind   int
	values []int
	bid    int
}

func (h *Hand) FillValues(cards string, round2 bool) {
	for i := 0; i < len(cards); i++ {
		card := string(cards[i])
		cardIndex := GetCardValue(card, round2)
		h.values = append(h.values, cardIndex)
	}
}

func (h *Hand) CalculateKind(part2 bool) {
	var countMap []int
	countMap = append(countMap, h.values...)
	sort.Slice(countMap, func(i, j int) bool {
		return countMap[i] > countMap[j]
	})
	buffer := 0
	lastValue := 0
	var groups []int
	for i := 0; i < len(countMap); i++ {
		if lastValue == countMap[i] || (part2 && countMap[i] == 1) {
			buffer += 1
			if i == len(countMap)-1 {
				groups = append(groups, buffer)
			}
			continue
		} else if buffer != 0 {
			groups = append(groups, buffer)
			buffer = 0
			lastValue = 0
		}
		buffer += 1
		lastValue = countMap[i]
		if i == len(countMap)-1 {
			groups = append(groups, buffer)
		}
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i] > groups[j]
	})
	if groups[0] == 5 {
		h.kind = 7
	} else if groups[0] == 4 {
		h.kind = 6
	} else if (groups[0] == 3 && groups[1] == 2) || (groups[0] == 2 && groups[1] == 3) {
		h.kind = 5
	} else if groups[0] == 3 && groups[1] == 1 {
		h.kind = 4
	} else if groups[0] == 2 && groups[1] == 2 {
		h.kind = 3
	} else if groups[0] == 2 && groups[1] == 1 {
		h.kind = 2
	} else if groups[0] == 1 {
		h.kind = 1
	}
}

func GetCardValue(card string, round2 bool) int {
	switch card {
	case "T":
		return 10
	case "J":
		if round2 {
			return 1
		}
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		cardInt, _ := strconv.Atoi(card)
		return cardInt
	}
}

func main() {
	b, err := os.ReadFile("day7/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")
	var hands []*Hand
	for i := 0; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], " ")
		bid, _ := strconv.Atoi(splitLine[1])
		hand := &Hand{bid: bid}
		hand.FillValues(splitLine[0], false)
		hand.CalculateKind(false)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].kind > hands[j].kind {
			return true
		} else if hands[i].kind == hands[j].kind {
			for k := 0; k < 5; k++ {
				if hands[i].values[k] == hands[j].values[k] {
					continue
				}
				if hands[i].values[k] > hands[j].values[k] {
					return true
				} else {
					return false
				}
			}
		}
		return false
	})

	puzzleSolution := 0
	slices.Reverse(hands)
	for i := 0; i < len(hands); i++ {
		bid := hands[i].bid
		multiplier := i + 1
		puzzleSolution += bid * multiplier
	}
	fmt.Printf("[Part1] Solution is %d\n", puzzleSolution)

	var hands2 []*Hand
	for i := 0; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], " ")
		bid, _ := strconv.Atoi(splitLine[1])
		hand := &Hand{bid: bid}
		hand.FillValues(splitLine[0], true)
		hand.CalculateKind(true)
		hands2 = append(hands2, hand)
	}

	sort.Slice(hands2, func(i, j int) bool {
		if hands2[i].kind > hands2[j].kind {
			return true
		} else if hands2[i].kind == hands2[j].kind {
			for k := 0; k < 5; k++ {
				if hands2[i].values[k] == hands2[j].values[k] {
					continue
				}
				if hands2[i].values[k] > hands2[j].values[k] {
					return true
				} else {
					return false
				}
			}
		}
		return false
	})

	puzzleSolution = 0
	slices.Reverse(hands2)
	for i := 0; i < len(hands2); i++ {
		bid := hands2[i].bid
		multiplier := i + 1
		puzzleSolution += bid * multiplier
	}
	fmt.Printf("[Part2] Solution is %d\n", puzzleSolution)

}
