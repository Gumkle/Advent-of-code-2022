package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	sums := make([]int, 0)
	var tmpSum int
	for {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			sums = append(sums, tmpSum)
			tmpSum = 0
			if err == io.EOF {
				break
			}
		}
		tmpSum += number
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println(sums[0])
	fmt.Println(sums[0] + sums[1] + sums[2])
}
