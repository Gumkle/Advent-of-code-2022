package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//func powerFor(ticks []int, pos int) int {
//	return pos * ticks[pos-1]
//}

func main() {
	input := bufio.NewReader(os.Stdin)
	currentX := 1
	ticks := make([]int, 1)
	ticks[0] = currentX
	for {
		line, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		argv := strings.Split(string(line), " ")
		if argv[0] == "noop" {
			ticks = append(ticks, currentX)
			continue
		}
		value, err := strconv.Atoi(argv[1])
		if err != nil {
			log.Fatalln(err)
		}
		ticks = append(ticks, currentX, currentX+value)
		currentX += value
	}

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			if math.Abs(float64(ticks[row*40+col]-col)) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
