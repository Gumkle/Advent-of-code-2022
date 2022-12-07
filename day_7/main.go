package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const cd = "$ cd"
const ls = "$ ls"
const treshold = 100000
const totalSpace = 70000000
const neededSpace = 30000000

type dir struct {
	parent   *dir
	children []*dir
	size     int
}

func (d *dir) resize(amount int) {
	d.size += amount
	if d.parent != nil {
		d.parent.resize(amount)
	}
}

func (d *dir) collectSizes() []int {
	sizes := []int{d.size}
	for _, child := range d.children {
		sizes = append(sizes, child.collectSizes()...)
	}
	return sizes
}

func filterSizes(sizes []int) []int {
	filteredSizes := make([]int, 0)
	for _, v := range sizes {
		if v <= treshold {
			filteredSizes = append(filteredSizes, v)
		}
	}
	return filteredSizes
}

func main() {
	input := bufio.NewReader(os.Stdin)
	var currentDir *dir
	for {
		byteLine, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		line := string(byteLine)
		if line == "$ cd .." {
			currentDir = currentDir.parent
			continue
		}
		if strings.HasPrefix(line, cd) {
			if currentDir == nil {
				currentDir = &dir{
					children: make([]*dir, 0),
				}
				continue
			}
			childDir := &dir{
				parent:   currentDir,
				children: make([]*dir, 0),
			}
			currentDir.children = append(currentDir.children, childDir)
			currentDir = childDir
			continue
		}
		if strings.HasPrefix(line, ls) {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			continue
		}
		resize, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		currentDir.resize(resize)
	}

	root := currentDir
	for {
		root = root.parent
		if root.parent == nil {
			break
		}
	}

	sizes := root.collectSizes()
	spaceOccupied := sizes[0]
	needToFreeUp := neededSpace - (totalSpace - spaceOccupied)
	//fmt.Println(needToFreeUp, spaceOccupied)
	sort.Ints(sizes)

	for _, v := range sizes {
		if v >= needToFreeUp {
			fmt.Println(v)
			break
		}
	}
	//belowTreshold := filterSizes(sizes)
	//var sum int
	//for _, v := range belowTreshold {
	//	sum += v
	//}
	//fmt.Println(sum)
}
