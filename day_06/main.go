package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const groupSize = 14

func main() {
	input := bufio.NewReader(os.Stdin)
	group := make([]byte, groupSize)

	_, err := io.ReadFull(input, group)
	if err != nil {
		log.Fatalln(err)
	}

	examinedChars := groupSize
	if !containsDuplicates(group) {
		fmt.Println(examinedChars)
		return
	}

	for {
		examinedChars++
		newByte, err := input.ReadByte()
		if err != nil {
			log.Fatalln(err)
		}
		group[(examinedChars-1)%groupSize] = newByte
		if !containsDuplicates(group) {
			fmt.Println(examinedChars)
			return
		}
	}
}

func containsDuplicates(group []byte) bool {
	hm := make(map[byte]struct{})
	for _, b := range group {
		if _, ok := hm[b]; ok {
			fmt.Println(string(group))
			return true
		}
		hm[b] = struct{}{}
	}
	return false
}
