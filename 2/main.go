package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  strategyGuide := bufio.NewScanner(input)

  totalScore := 0

  for strategyGuide.Scan() {
    strategy := strategyGuide.Text()

    if strategy == "" {
      break
    }

    opponentChoice, myChoice := strings.Split(strategy, " ")[0], strings.Split(strategy, " ")[1]

    switch {
    case opponentChoice == "A" && myChoice == "X":
      totalScore += 4
    case opponentChoice == "A" && myChoice == "Y":
      totalScore += 8
    case opponentChoice == "A" && myChoice == "Z":
      totalScore += 3
    case opponentChoice == "B" && myChoice == "X":
      totalScore += 1
    case opponentChoice == "B" && myChoice == "Y":
      totalScore += 5
    case opponentChoice == "B" && myChoice == "Z":
      totalScore += 9
    case opponentChoice == "C" && myChoice == "X":
      totalScore += 7
    case opponentChoice == "C" && myChoice == "Y":
      totalScore += 2
    case opponentChoice == "C" && myChoice == "Z":
      totalScore += 6
    }

  }

  fmt.Println(totalScore)

}
