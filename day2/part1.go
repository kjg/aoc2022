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
	fmt.Println("My total score: ", totalScore)
}

func roundScore(round []string) int {
	return shapeScore(round[1]) + outcomeScore(round[0], round[1])
}

func shapeScore(me string) int {
	shapeScores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	return shapeScores[me]
}

func outcomeScore(opponent, me string) int {

	opponentShape := shapeLookup(opponent)
	myShape := shapeLookup(me)

	if myShape == opponentShape {
		return 3
	}
	if myShape == "Rock" && opponentShape == "Scissors" {
		return 6
	}
	if myShape == "Paper" && opponentShape == "Rock" {
		return 6
	}
	if myShape == "Scissors" && opponentShape == "Paper" {
		return 6
	}
	return 0
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
