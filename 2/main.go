package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  score := part1()
  fmt.Println(score)

  newScore := part2()
  fmt.Println(newScore)
}


func part1() int {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  strategyGuide := bufio.NewScanner(input)

  var score int
	scores := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4,"B Y": 5,"C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}

  for strategyGuide.Scan() {
    strategy := strategyGuide.Text()
    score += scores[strategy]
  }

  return score
}

func part2() int {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  strategyGuide := bufio.NewScanner(input)

  var score int
	scores := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4,"B Y": 5,"C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}
  for strategyGuide.Scan() {
    strategy := strategyGuide.Text()
    score += scores[strategy]
  }

  return score
}
