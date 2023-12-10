package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	rangeStart int
	rangeEnd   int
	offset     int
}

func Parse(input string) Mapping {
	valueSlice := strings.Split(input, " ")

	mappedValue, _ := strconv.Atoi(valueSlice[0])
	start, _ := strconv.Atoi(valueSlice[1])
	rangeLength, _ := strconv.Atoi(valueSlice[2])

	return Mapping{
		rangeStart: start,
		rangeEnd:   start + rangeLength - 1,
		offset:     mappedValue - start,
	}
}

type Maps struct {
	value []Mapping
}

func NewMaps(block string) *Maps {
	lines := strings.Split(block, "\n")[1:]
	var mappings []Mapping
	for i := 0; i < len(lines); i++ {
		mappings = append(mappings, Parse(lines[i]))
	}
	return &Maps{value: mappings}
}

func (m *Maps) GetMappedValue(input int) int {
	var mappedValue *int
	for i := 0; i < len(m.value); i++ {
		currentMapping := m.value[i]
		if input >= currentMapping.rangeStart && input <= currentMapping.rangeEnd {
			mappedValue = new(int)
			*mappedValue = input + currentMapping.offset
			break
		}

	}
	if mappedValue != nil {
		return *mappedValue
	} else {
		return input
	}
}

func (m *Maps) GetMappedValue2(input []int) [][]int {
	var mappedValues [][]int
	for i := 0; i < len(m.value); i++ {
		currentMapping := m.value[i]
		overlapStart, overlapEnd := CalculateOverlap(input[0], input[1], currentMapping.rangeStart, currentMapping.rangeEnd)
		if overlapStart == -1 && overlapEnd == -1 {
			continue
		}
		mappedValues = append(mappedValues, []int{overlapStart + currentMapping.offset, overlapEnd + currentMapping.offset})

	}
	if len(mappedValues) == 0 {
		mappedValues = append(mappedValues, input)
	}

	return mappedValues
}

func main() {
	b, err := os.ReadFile("day5/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	blocks := strings.Split(string(b), "\n\n")

	seeds := strings.Split(blocks[0][7:], " ")

	correspondingKey, smallestValue := CalculateMinLocation(blocks, seeds)

	fmt.Printf("[Part1] Seed: %s with Location %d\n", correspondingKey, smallestValue)

	var maps []*Maps

	for _, block := range blocks[1:] {
		createdMap := NewMaps(block)
		maps = append(maps, createdMap)
	}

	var seeds2 [][]int

	for i := 0; i < len(seeds); i += 2 {
		rangeStart, _ := strconv.Atoi(seeds[i])
		rangeLength, _ := strconv.Atoi(seeds[i+1])
		value := []int{rangeStart, rangeStart + rangeLength - 1}
		seeds2 = append(seeds2, value)
	}

	locationRange := CalculateMinLocation2(blocks, seeds2)

	fmt.Printf("[Part2] Location %d\n", locationRange[0][0])

}

func CalculateOverlap(seedsStart int, seedEnd int, mappingStart int, mappingEnd int) (int, int) {
	if mappingStart > seedEnd || seedsStart > mappingEnd {
		return -1, -1
	}

	start := max(seedsStart, mappingStart)
	end := min(seedEnd, mappingEnd)

	return start, end
}

func CalculateMinLocation2(blocks []string, seeds [][]int) [][]int {

	var maps []*Maps

	for _, block := range blocks[1:] {
		createdMap := NewMaps(block)
		maps = append(maps, createdMap)
	}

	for i := 0; i < len(maps); i++ {
		currentMap := maps[i]
		tmp := make([][]int, 0)
		for j := 0; j < len(seeds); j++ {
			tmp = append(tmp, currentMap.GetMappedValue2(seeds[j])...)
		}
		if tmp != nil {
			seeds = tmp
		}
	}

	sortSlice(seeds)

	return seeds
}

func sortSlice(sl [][]int) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][0] < sl[j][0]
	})
}

func CalculateMinLocation(blocks []string, seeds []string) (string, int) {
	var maps []*Maps

	for _, block := range blocks[1:] {
		createdMap := NewMaps(block)
		maps = append(maps, createdMap)
	}

	seedToLocationMap := make(map[string]int)

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		returnedValue := maps[0].GetMappedValue(seedInt)

		for i := 1; i < len(maps); i++ {
			currentMap := maps[i]
			returnedValue = currentMap.GetMappedValue(returnedValue)
		}
		seedToLocationMap[seed] = returnedValue
	}

	smallestValue := math.MaxInt64
	correspondingKey := ""
	for key, value := range seedToLocationMap {
		if value < smallestValue {
			smallestValue = value
			correspondingKey = key
		}
	}

	return correspondingKey, smallestValue
}
