package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  sumOfPriorities := part1()
  fmt.Println(sumOfPriorities)

  p := part2()
  fmt.Println(p)

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

func part2() int {
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
    firstRucksack := rucksackList.Text()
    itemsFirstRucksack := createItemsCheck(firstRucksack)

    rucksackList.Scan()
    secondRucksack := rucksackList.Text()
    itemsSecondRucksack := createItemsCheck(secondRucksack)

    rucksackList.Scan()
    thirdRucksack := rucksackList.Text()
    itemsThirdRucksack := createItemsCheck(thirdRucksack)


    for item := range itemsFirstRucksack {
      if itemsSecondRucksack[item] && itemsThirdRucksack[item] {
        sumOfPriorities += typePriority[item]
      }
    }

  }

  return sumOfPriorities

}

func createItemsCheck (items string) (itemsCheck map[rune]bool) {
  itemsCheck = make(map[rune]bool)

  for _, item := range items {
    itemsCheck[item] = true
  }
  return
}
