package main

import (
  "bufio"
  "fmt"
  "log"
  "strconv"
  "os"
)

func main() {
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

  var mostCalories, largestElf int
  for index, element := range elfCalories {
    if element > mostCalories {
      mostCalories = element
      largestElf = index
    }
  }
  fmt.Println("Elf", largestElf, "has the most calories:", mostCalories)

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}
