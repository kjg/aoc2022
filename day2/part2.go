package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScore = 0

	for scanner.Scan() {
		roundTxt := scanner.Text()
		round := strings.Split(roundTxt, " ")
		roundScore := part2RoundScore(round)
		totalScore += roundScore
	}
	fmt.Println("Part2 total score: ", totalScore)
}

var desiredOutcomes = map[string]string{
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

func part2RoundScore(round []string) int {
	desiredOutcome := desiredOutcomes[round[1]]
	myShape := myNeededShape(round[0], desiredOutcome)
	return shapeScore(myShape) + outcomeScore(desiredOutcome)
}

func myNeededShape(opponentShape, desiredOutcome string) string {
	opponentShape = shapeLookup(opponentShape)
	if desiredOutcome == "Draw" {
		return opponentShape
	}
	if desiredOutcome == "Win" {
		return whatBeats(opponentShape)
	}
	return shapeFails(opponentShape)
}

var shapeBeats = map[string]string{
	"Rock":     "Paper",
	"Paper":    "Scissors",
	"Scissors": "Rock",
}

func whatBeats(shape string) string {
	return shapeBeats[shape]
}

func shapeFails(shape string) string {
	return shapeBeats[whatBeats(shape)]
}
