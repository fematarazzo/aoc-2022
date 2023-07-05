package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  maxCalories := part1()
  fmt.Println(maxCalories)  

  totalCalories := part2()
  fmt.Println(totalCalories)  

}

func part1() int {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  caloriesLog := bufio.NewScanner(input)
  maxCalories, currentCalories := 0, 0

  for caloriesLog.Scan() {
    calories, err := strconv.Atoi(caloriesLog.Text())
    currentCalories += calories

    if err != nil {

      if currentCalories >= maxCalories {
        maxCalories = currentCalories
      }

      currentCalories = 0

    }
  }

  return maxCalories

}

func part2() int {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  caloriesLog := bufio.NewScanner(input)
  firstPlace, secondPlace, thirdPlace, currentCalories := 0, 0, 0, 0

  for caloriesLog.Scan() {
    calories, err := strconv.Atoi(caloriesLog.Text())
    currentCalories += calories

    if err != nil {
      switch {
      case currentCalories > thirdPlace:
        thirdPlace = currentCalories 
      case thirdPlace > secondPlace:
        thirdPlace, secondPlace = secondPlace, thirdPlace
      case secondPlace > firstPlace:
        secondPlace, firstPlace = firstPlace, secondPlace
      }

      currentCalories = 0

    }
  }

  totalCalories := firstPlace + secondPlace + thirdPlace

  return totalCalories

}
