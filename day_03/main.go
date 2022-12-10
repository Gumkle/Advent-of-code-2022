package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	//"strings"
)

const lowercaseOffset = -96
const uppercaseOffset = 58
const groupSize = 3

func calcPriority(in rune) int {
	res := int(in) + lowercaseOffset
	if res > 0 {
		return res
	}
	return res + uppercaseOffset
}

func main() {
	var sum int
read:
	for {
		var group [groupSize]string
		for i := 0; i < groupSize; i++ {
			_, err := fmt.Scanf("%s", &group[i])
			if err != nil {
				if err == io.EOF {
					break read
				}
				log.Fatalln(err)
			}
		}

		for _, c := range group[0] {
			thereItIs := true
			for i := 1; i < groupSize; i++ {
				if !strings.ContainsRune(group[i], c) {
					thereItIs = false
				}
			}
			if thereItIs {
				sum += calcPriority(c)
				break
			}
		}
	}
	fmt.Println(sum)
}
