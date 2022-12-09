package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	headPosition := GridPosition{0, 0}
	tailPosition := GridPosition{0, 0}

	tailVisits := make(map[GridPosition]bool)
	tailVisits[tailPosition] = true

	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	moves := strings.Split(stream, "\n")
	moves = moves[0 : len(moves)-1]

	for _, move := range moves {
		direction := string(move[0])
		distance, _ := strconv.Atoi(string(move[2:]))

		for i := 0; i < distance; i++ {
			headPosition = moveHead(headPosition, direction)
			tailPosition = catchUpTail(headPosition, tailPosition)
			tailVisits[tailPosition] = true
		}
	}
	fmt.Println("The tail visits", len(tailVisits), "unique locations.")
}

type GridPosition struct {
	x int
	y int
}

func moveHead(headPosition GridPosition, direction string) GridPosition {
	switch direction {
	case "U":
		headPosition.y++
	case "D":
		headPosition.y--
	case "R":
		headPosition.x++
	case "L":
		headPosition.x--
	}
	return headPosition
}

func catchUpTail(headPosition GridPosition, tailPosition GridPosition) GridPosition {
	xDelta := headPosition.x - tailPosition.x
	yDelta := headPosition.y - tailPosition.y

	if xDelta > 1 {
		tailPosition.x = headPosition.x - 1
		tailPosition.y = headPosition.y
	} else if xDelta < -1 {
		tailPosition.x = headPosition.x + 1
		tailPosition.y = headPosition.y
	} else if yDelta > 1 {
		tailPosition.y = headPosition.y - 1
		tailPosition.x = headPosition.x
	} else if yDelta < -1 {
		tailPosition.y = headPosition.y + 1
		tailPosition.x = headPosition.x
	}
	return tailPosition
}
