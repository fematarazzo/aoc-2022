package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

  fmt.Println(score)

}
