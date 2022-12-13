package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const y = 1
const x = 0

type node struct {
	pos     [2]int
	value   int
	visited bool
}

func main() {
	input := bufio.NewReader(os.Stdin)
	fieldMap := make([][]rune, 0)
	for {
		line, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		fieldMap = append(fieldMap, []rune(string(line)))
	}

	//var startPos [2]int
	startPoses := make([][2]int, 0)
	var endPos [2]int
	for y, line := range fieldMap {
		for x, letter := range line {
			//if letter == 'S' {
			//	startPos = [2]int{x, y}
			//}
			if letter == 'S' || letter == 'a' {
				startPoses = append(startPoses, [2]int{x, y})
			}
			if letter == 'E' {
				endPos = [2]int{x, y}
			}
		}
	}

	foundLens := make([]int, 0)
	for _, startPos := range startPoses {
		unvisitedNodes := make([]*node, 0)
		unvisitedNodes = append(unvisitedNodes, &node{pos: startPos, value: 0})
		visited := make(map[[2]int]*node, 0)
		foundLens = append(foundLens, wayLenTo(endPos, unvisitedNodes, fieldMap, visited))
	}
	sort.Ints(foundLens)
	fmt.Println(foundLens)
}

func wayLenTo(end [2]int, unvisited []*node, fieldMap [][]rune, nodeRegistry map[[2]int]*node) int {
	if len(unvisited) == 0 {
		return 0
	}
	current := unvisited[0]
	nodeRegistry[current.pos] = current
	if len(unvisited) == 1 {
		unvisited = []*node{}
	} else {
		unvisited = unvisited[1:]
	}
	if current.visited {
		return wayLenTo(end, unvisited, fieldMap, nodeRegistry)
	}
	if current.pos == end {
		return current.value
	}
	neighbours := findNeighboursFor(current.pos, fieldMap)
	unvisitedNeighbours := make([]*node, 0)
	for _, pos := range neighbours {
		if neighbour, ok := nodeRegistry[pos]; !ok {
			unvisitedNeighbours = append(unvisitedNeighbours, &node{
				pos:     pos,
				value:   999,
				visited: false,
			})
		} else {
			if !neighbour.visited {
				unvisitedNeighbours = append(unvisitedNeighbours, neighbour)
			}
		}
	}
	if len(unvisitedNeighbours) == 0 {
		current.visited = true
		return wayLenTo(end, unvisited, fieldMap, nodeRegistry)
	}

	for _, neighbour := range unvisitedNeighbours {
		neighbour.value = current.value + 1
		nodeRegistry[neighbour.pos] = neighbour
		unvisited = append(unvisited, neighbour)
	}
	current.visited = true
	return wayLenTo(end, unvisited, fieldMap, nodeRegistry)
}

func findNeighboursFor(pos [2]int, fieldMap [][]rune) [][2]int {
	currentSymbol := fieldMap[pos[y]][pos[x]]
	result := make([][2]int, 0)
	nextPos := [2]int{pos[x], pos[y] - 1}
	if pos[y] > 0 {
		nextSymbol := fieldMap[nextPos[y]][nextPos[x]]
		if canWalkTo(currentSymbol, nextSymbol) {
			result = append(result, nextPos)
		}
	}

	nextPos = [2]int{pos[x], pos[y] + 1}
	if pos[y] < len(fieldMap)-1 {
		nextSymbol := fieldMap[nextPos[y]][nextPos[x]]
		if canWalkTo(currentSymbol, nextSymbol) {
			result = append(result, nextPos)
		}
	}

	nextPos = [2]int{pos[x] - 1, pos[y]}
	if pos[x] > 0 {
		nextSymbol := fieldMap[nextPos[y]][nextPos[x]]
		if canWalkTo(currentSymbol, nextSymbol) {
			result = append(result, nextPos)
		}
	}

	nextPos = [2]int{pos[x] + 1, pos[y]}
	if pos[x] < len(fieldMap[0])-1 {
		nextSymbol := fieldMap[nextPos[y]][nextPos[x]]
		if canWalkTo(currentSymbol, nextSymbol) {
			result = append(result, nextPos)
		}
	}

	return result
}

func canWalkTo(currentSymbol, nextSymbol rune) bool {
	if nextSymbol == 'E' {
		nextSymbol = 'z'
	}
	if currentSymbol == 'S' {
		currentSymbol = 'a'
	}
	return nextSymbol <= (currentSymbol + 1)
}
