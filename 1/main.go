package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  input, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
    return
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

  fmt.Println(maxCalories)

}
