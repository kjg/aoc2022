package main

import (
  "bufio"
  "fmt"
  "log"
  "strconv"
  "os"
  "sort"
)

func Part2() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var elfNumber = 0
  var elfCalories = []int{0}

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    text := scanner.Text()
    if(text == "") {
      elfNumber++
      elfCalories = append(elfCalories, 0)
    } else {
      calories, _ := strconv.Atoi(text)
      elfCalories[elfNumber] += calories
    }
  }

  sort.Ints(elfCalories)

  threeLargest := elfCalories[len(elfCalories)-3:]
  var totalCalories = 0
  for _, element := range threeLargest {
    totalCalories += element
  }
  fmt.Println("The total of the largest 3:", totalCalories)

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}
