package main

import (
	"fmt"
	"io"
	"log"
)

const (
	win  = 6
	draw = 3
	lost = 0
)

const (
	theirPaper    = "B"
	theirRock     = "A"
	theirScissors = "C"
)

const (
	ourPaper    = "Y"
	ourRock     = "X"
	ourScissors = "Z"
)

const (
	resultWin  = "Z"
	resultDraw = "Y"
	resultLose = "X"
)

var resultPoints = map[string]int{
	resultLose: lost,
	resultDraw: draw,
	resultWin:  win,
}

var figurePoints = map[string]int{
	ourPaper:    2,
	ourRock:     1,
	ourScissors: 3,
}

var conclusion = map[string]map[string]int{
	theirScissors: {
		ourPaper:    lost,
		ourRock:     win,
		ourScissors: draw,
	},
	theirPaper: {
		ourScissors: win,
		ourRock:     lost,
		ourPaper:    draw,
	},
	theirRock: {
		ourScissors: lost,
		ourRock:     draw,
		ourPaper:    win,
	},
}

var responses = map[string]map[string]int{
	theirRock: {
		resultWin:  figurePoints[ourPaper],
		resultLose: figurePoints[ourScissors],
		resultDraw: figurePoints[ourRock],
	},
	theirPaper: {
		resultWin:  figurePoints[ourScissors],
		resultDraw: figurePoints[ourPaper],
		resultLose: figurePoints[ourRock],
	},
	theirScissors: {
		resultWin:  figurePoints[ourRock],
		resultDraw: figurePoints[ourScissors],
		resultLose: figurePoints[ourPaper],
	},
}

func main() {
	var pointsTotal int
	for {
		var mine, theirs string
		_, err := fmt.Scanf("%s %s", &theirs, &mine)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		//pointsTotal += figurePoints[mine] + conclusion[theirs][mine]
		pointsTotal += resultPoints[mine] + responses[theirs][mine]
	}
	fmt.Println(pointsTotal)
}
