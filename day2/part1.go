package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() {
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
		roundScore := roundScore(round)
		totalScore += roundScore
	}
	fmt.Println("Part1 total score: ", totalScore)
}

func roundScore(round []string) int {
	return shapeScore(shapeLookup(round[1])) + outcomeScore(outcome(round[0], round[1]))
}

func shapeScore(me string) int {
	shapeScores := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}
	return shapeScores[me]
}

func outcomeScore(outcome string) int {
	outcomeScores := map[string]int{
		"Win":  6,
		"Lose": 0,
		"Draw": 3,
	}
	return outcomeScores[outcome]
}

func outcome(opponent, me string) string {

	opponentShape := shapeLookup(opponent)
	myShape := shapeLookup(me)

	if myShape == opponentShape {
		return "Draw"
	}
	if myShape == "Rock" && opponentShape == "Scissors" {
		return "Win"
	}
	if myShape == "Paper" && opponentShape == "Rock" {
		return "Win"
	}
	if myShape == "Scissors" && opponentShape == "Paper" {
		return "Win"
	}
	return "Lose"
}

func shapeLookup(shape string) string {
	shapeLookup := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	return shapeLookup[shape]
}
