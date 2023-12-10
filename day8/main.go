package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	current string
	left    string
	right   string
}

var nodes = map[string]*Node{}

func main() {
	b, err := os.ReadFile("day8/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")

	instructions := lines[0]

	var startNodes []string

	for _, nodeString := range lines[2:] {
		current := nodeString[:3]
		if current[2:3] == "A" {
			startNodes = append(startNodes, current)
		}
		left := nodeString[strings.IndexByte(nodeString, '(')+1 : strings.IndexByte(nodeString, ',')]
		right := nodeString[strings.IndexByte(nodeString, ',')+2 : strings.IndexByte(nodeString, ')')]
		nodes[current] = &Node{current: current, left: left, right: right}
	}

	stepsNeeded := TraverseNodes(instructions, "AAA", 1)

	fmt.Printf("[Part1] Total Steps needed: %d\n", stepsNeeded)

	var paths []int

	for _, node := range startNodes {
		paths = append(paths, TraverseNodes(instructions, node, 1))
	}
	stepsNeeded = LCM(paths[0], paths[1], paths...)

	fmt.Printf("[Part2] Total Steps needed: %d\n", stepsNeeded)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func TraverseNodes(instructions string, currentNode string, counter int) int {
	for i := 0; i < len(instructions); i++ {
		direction := string(instructions[i])
		if direction == "R" {
			currentNode = nodes[currentNode].right
		} else {
			currentNode = nodes[currentNode].left
		}
		if currentNode[2:3] == "Z" {
			return counter
		}
		counter += 1
	}

	return TraverseNodes(instructions, currentNode, counter)
}
