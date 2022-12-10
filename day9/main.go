package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	moves := strings.Split(stream, "\n")
	moves = moves[0 : len(moves)-1]

	moveRope(makeRope(2), moves)
	moveRope(makeRope(10), moves)
}

type GridPosition struct {
	x int
	y int
}

func makeRope(length int) []GridPosition {
	rope := make([]GridPosition, length)
	for i := 0; i < length; i++ {
		rope[i] = GridPosition{0, 0}
	}
	return rope
}

func moveRope(rope []GridPosition, moves []string) {
	tailVisits := make(map[GridPosition]bool)
	tailVisits[GridPosition{0, 0}] = true

	for _, move := range moves {
		direction := string(move[0])
		distance, _ := strconv.Atoi(string(move[2:]))

		for i := 0; i < distance; i++ {
			rope[0] = moveHead(rope[0], direction)
			for i := 1; i < len(rope); i++ {
				rope[i] = catchUpTail(rope[i-1], rope[i])
			}
			tailVisits[rope[len(rope)-1]] = true
		}
	}
	fmt.Println("The tail visits", len(tailVisits), "unique locations.")
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

	xMoved := false
	yMoved := false

	if xDelta > 1 {
		tailPosition.x = headPosition.x - 1
		xMoved = true
	} else if xDelta < -1 {
		tailPosition.x = headPosition.x + 1
		xMoved = true
	}

	if yDelta > 1 {
		tailPosition.y = headPosition.y - 1
		yMoved = true
	} else if yDelta < -1 {
		tailPosition.y = headPosition.y + 1
		yMoved = true
	}

	if xMoved && !yMoved {
		tailPosition.y = headPosition.y
	}

	if yMoved && !xMoved {
		tailPosition.x = headPosition.x
	}

	return tailPosition
}
