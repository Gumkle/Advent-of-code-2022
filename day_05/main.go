package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type command struct {
	from  int
	to    int
	quant int
}

func main() {
	input := bufio.NewReader(os.Stdin)
	stacks := make(map[int][]rune, 0)
	for {
		line, err := input.ReadBytes(byte('\n'))
		if err != nil {
			log.Fatalln(err)
		}
		if string(line) == "\n" {
			break
		}
		
	}

	var quant, src, dst int
	commandQueue := make([]*command, 0)
	for {
		_, err := fmt.Fscanf(input, "move %d from %d to %d\n", &quant, &src, &dst)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			}
			log.Fatalln(err)
		}
		commandQueue = append(commandQueue, &command{
			from:  src,
			to:    dst,
			quant: quant,
		})
	}
}
