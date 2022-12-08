package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	treeMap := make([][]int, 0)
	for {
		byteline, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		line := string(byteline)
		mapRow := make([]int, 0)
		for _, v := range line {
			if v == '\n' {
				break
			}
			mapRow = append(mapRow, int(v-'0'))
		}
		treeMap = append(treeMap, mapRow)
	}
	//perimeterVisible := (len(treeMap)-2)*2 + len(treeMap[0])*2
	//visible := perimeterVisible
	scores := make([]int, 0)
	for y := 1; y < len(treeMap)-1; y++ {
		for x := 1; x < len(treeMap[0])-1; x++ {
			var topScore, botScore, leftScore, rightScore int
			fmt.Printf("consider %dx %dy (%d)\n", x, y, treeMap[y][x])

			for i := x - 1; i >= 0; i-- {
				if treeMap[y][i] >= treeMap[y][x] {
					break
				}
				leftScore++
			}
			fmt.Println("left:", leftScore)
			for i := x + 1; i < len(treeMap[0]); i++ {
				rightScore++
				if treeMap[y][i] >= treeMap[y][x] {
					break
				}
			}
			fmt.Println("right:", rightScore)
			for i := y - 1; i >= 0; i-- {
				topScore++
				if treeMap[i][x] >= treeMap[y][x] {
					break
				}
			}
			fmt.Println("top:", topScore)
			for i := y + 1; i < len(treeMap); i++ {
				botScore++
				if treeMap[i][x] >= treeMap[y][x] {
					break
				}
			}
			fmt.Println("bot:", botScore)

			scores = append(scores, topScore*botScore*leftScore*rightScore)

			//upper := make([]int, 0)
			//lower := make([]int, 0)
			//right := make([]int, 0)
			//left := make([]int, 0)
			//for i := 0; i < x; i++ {
			//	left = append(left, treeMap[y][i])
			//}
			//for i := x + 1; i < len(treeMap[0]); i++ {
			//	right = append(right, treeMap[y][i])
			//}
			//for i := 0; i < y; i++ {
			//	upper = append(upper, treeMap[i][x])
			//}
			//for i := y + 1; i < len(treeMap); i++ {
			//	lower = append(lower, treeMap[i][x])
			//}
			//sort.Ints(left)
			//sort.Ints(right)
			//sort.Ints(upper)
			//sort.Ints(lower)
			////fmt.Printf("for %dx %dy (%d): ", x, y, treeMap[y][x])
			////fmt.Println(left, right, upper, lower)
			//if left[len(left)-1] < treeMap[y][x] {
			//	visible++
			//	continue
			//}
			//if right[len(right)-1] < treeMap[y][x] {
			//	visible++
			//	continue
			//}
			//if upper[len(upper)-1] < treeMap[y][x] {
			//	visible++
			//	continue
			//}
			//if lower[len(lower)-1] < treeMap[y][x] {
			//	visible++
			//	continue
			//}
		}
	}
	sort.Ints(scores)
	fmt.Println(scores)
}
