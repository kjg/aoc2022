package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)
	rows := strings.Split(stream, "\n")
	rows = rows[:len(rows)-1]

	var startPos Coordinate
	var endPos Coordinate

	elevations := make(map[Coordinate]string)
	height := len(rows)
	width := len(rows[0])

	for rowNum, row := range rows {
		for colNum, col := range row {
			if string(col) == "S" {
				startPos = Coordinate{colNum, rowNum}
			}
			if string(col) == "E" {
				endPos = Coordinate{colNum, rowNum}
			}

			elevations[Coordinate{colNum, rowNum}] = string(col)
		}
	}

	heightmap := Heightmap{startPos, endPos, make(map[Coordinate]bool), make([]Coordinate, 0, height*width), make(map[Coordinate]int), elevations, height, width}

	heightmap.toDo = append(heightmap.toDo, endPos)
	heightmap.doSearch()

	fmt.Println(heightmap.distances[startPos])
}

func (h Heightmap) doSearch() {

	for len(h.toDo) > 0 {
		check := h.toDo[0]
		if check == h.endPos {
			h.distances[check] = 0
		}

		above := Coordinate{check.col, check.row - 1}
		below := Coordinate{check.col, check.row + 1}
		left := Coordinate{check.col - 1, check.row}
		right := Coordinate{check.col + 1, check.row}

		neighbors := []Coordinate{above, below, left, right}
		for _, neighbor := range neighbors {
			currentElevation := h.elevations[check]
			if currentElevation == "E" {
				currentElevation = "z"
			}
			if currentElevation == "S" {
				currentElevation = "a"
			}
			nElevation, hasElevation := h.elevations[neighbor]
			if hasElevation {
				if nElevation == "S" {
					nElevation = "a"
				}
				if nElevation == "E" {
					nElevation = "z"
				}

				fmt.Println("Neighbor", neighbor, nElevation, currentElevation)
				if h.processeed[neighbor] {
					fmt.Println("Already processed", neighbor)
					if currentElevation[0]+1 >= nElevation[0] {
						fmt.Println("And it is the parent")
						h.distances[check] = h.distances[neighbor] + 1
					}
				} else if nElevation[0]+1 >= currentElevation[0] {

					alreadyTodo := false
					for _, existing := range h.toDo {
						if existing == neighbor {
							alreadyTodo = true
						}
					}
					if !alreadyTodo {
						h.toDo = append(h.toDo, neighbor)
					}
				}
			}

		}

		h.processeed[check] = true
		if check == h.startPos {
			fmt.Println("Found start", check)
			h.toDo = []Coordinate{}
		} else {
			h.toDo = h.toDo[1:]
		}

		fmt.Println("checked", check, h.distances[check], "To do", h.toDo)
	}

}

type Heightmap struct {
	startPos   Coordinate
	endPos     Coordinate
	processeed map[Coordinate]bool
	toDo       []Coordinate
	distances  map[Coordinate]int
	elevations map[Coordinate]string
	height     int
	width      int
}

type Coordinate struct {
	col int
	row int
}
