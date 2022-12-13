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

	singleStart := make([]Coordinate, 1)
	startPositions := make([]Coordinate, 0)
	var endPos Coordinate

	elevations := make(map[Coordinate]string)
	height := len(rows)
	width := len(rows[0])

	for rowNum, row := range rows {
		for colNum, col := range row {
			if string(col) == "S" {
				singleStart[0] = Coordinate{colNum, rowNum}
				startPositions = append(startPositions, Coordinate{colNum, rowNum})
			}
			if string(col) == "a" {
				startPositions = append(startPositions, Coordinate{colNum, rowNum})
			}
			if string(col) == "E" {
				endPos = Coordinate{colNum, rowNum}
			}

			elevations[Coordinate{colNum, rowNum}] = string(col)
		}
	}

	heightmap := Heightmap{singleStart, endPos, make(map[Coordinate]bool), make([]Coordinate, 0, height*width), make(map[Coordinate]int), elevations, height, width}

	heightmap.toDo = append(heightmap.toDo, endPos)
	shortestPath := heightmap.doSearch()

	fmt.Println("The shortest path is", shortestPath, "steps long")

	heightmap = Heightmap{startPositions, endPos, make(map[Coordinate]bool), make([]Coordinate, 0, height*width), make(map[Coordinate]int), elevations, height, width}

	heightmap.toDo = append(heightmap.toDo, endPos)
	shortestPath = heightmap.doSearch()

	fmt.Println("The shortest path is", shortestPath, "steps long")
}

func (h Heightmap) doSearch() int {

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
		if len(h.startPositions) == 0 {
			fmt.Println("Found all starts")
			h.toDo = []Coordinate{}
		} else {
			h.toDo = h.toDo[1:]
		}

		fmt.Println("checked", check, h.distances[check], "To do", h.toDo)
	}

	shortestPath := -1
	for _, startPos := range h.startPositions {
		currentDistance := h.distances[startPos]
		if (currentDistance < shortestPath || shortestPath == -1) && currentDistance > 0 {
			shortestPath = h.distances[startPos]
		}
	}
	return shortestPath
}

type Heightmap struct {
	startPositions []Coordinate
	endPos         Coordinate
	processeed     map[Coordinate]bool
	toDo           []Coordinate
	distances      map[Coordinate]int
	elevations     map[Coordinate]string
	height         int
	width          int
}

type Coordinate struct {
	col int
	row int
}
