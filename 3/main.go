package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  sumOfPriorities := part1()
  fmt.Println(sumOfPriorities)

}

func part1() int {
  input, err := os.Open("input.txt")
  
  if err != nil {
    fmt.Println("Error opening file", err)
  }

  rucksackList := bufio.NewScanner(input)

  typePriority := make(map[rune]int)

  for i, letter := 0, 'a'; letter <= 'z'; i, letter = i + 1, letter + 1 {
    typePriority[letter] = i + 1
  }

  for i, letter := 26, 'A'; letter <= 'Z'; i, letter = i + 1, letter + 1 {
    typePriority[letter] = i + 1
  }

  sumOfPriorities := 0

  for rucksackList.Scan() {
    rucksack := rucksackList.Text()
    separator := len(rucksack) / 2
    items := make(map[rune]bool)
    firstCompartment := rucksack[:separator]
    secondCompartment := rucksack[separator:]

    for _, itemFirstComponent := range firstCompartment {
      items[itemFirstComponent] = true
    }

    for _, itemSecondComponent := range secondCompartment {
      if items[itemSecondComponent] {
        sumOfPriorities += typePriority[itemSecondComponent]
        break
      }
    }

  }

  return sumOfPriorities

}
