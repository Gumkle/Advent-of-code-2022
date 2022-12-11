package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type position struct {
	x int
	y int
}

func (p *position) move(dir direction) {
	if dir == left {
		p.x -= 1
	}
	if dir == right {
		p.x += 1
	}
	if dir == up {
		p.y += 1
	}
	if dir == down {
		p.y -= 1
	}
}

func (p *position) touches(p2 *position) bool {
	return math.Abs(float64(p.x-p2.x)) <= 1 && math.Abs(float64(p.y-p2.y)) <= 1
}

func (p *position) follow(p2 *position) {
	if p.touches(p2) {
		return
	}
	if p.x == p2.x {
		if p.y < p2.y {
			p.move(up)
		}
		if p.y > p2.y {
			p.move(down)
		}
	}
	if p.y == p2.y {
		if p.x < p2.x {
			p.move(right)
		}
		if p.x > p2.x {
			p.move(left)
		}
	}
	if p.x < p2.x && p.y < p2.y {
		p.move(up)
		p.move(right)
	}
	if p.x < p2.x && p.y > p2.y {
		p.move(down)
		p.move(right)
	}
	if p.x > p2.x && p.y < p2.y {
		p.move(up)
		p.move(left)
	}
	if p.x > p2.x && p.y > p2.y {
		p.move(down)
		p.move(left)
	}
}

type direction string

const (
	left  = direction("L")
	right = direction("R")
	up    = direction("U")
	down  = direction("D")
)

type command struct {
	dir direction
	len int
}

const howManyKnots = 9

func main() {
	input := bufio.NewReader(os.Stdin)
	positions := make(map[position]bool, 0)
	//Tpos := &position{}
	knots := make([]position, howManyKnots)
	Hpos := &position{}

	for {
		command := &command{}
		_, err := fmt.Fscanf(input, "%s %d\n", &command.dir, &command.len)
		if err != nil {
			break
		}
		for i := 0; i < command.len; i++ {
			Hpos.move(command.dir)
			knots[0].follow(Hpos)
			for i := 1; i < howManyKnots; i++ {
				knots[i].follow(&knots[i-1])
			}
			positions[knots[howManyKnots-1]] = true
		}
	}
	fmt.Println(len(positions))
}
