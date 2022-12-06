package main

import (
	"fmt"
	"io"
)

func main() {
	var count int
	for {
		var first, second [2]int
		_, err := fmt.Scanf("%d-%d,%d-%d", &first[0], &first[1], &second[0], &second[1])
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		//if first[0] <= second[0] && first[1] >= second[1] || second[0] <= first[0] && second[1] >= first[1] {
		//	count++
		//}
		if first[0] <= second[1] && second[0] <= first[1] {
			count++
		}
	}
	fmt.Println(count)
}
