package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "strings"
  "strconv"
)

func readMass() []int {
  // Read entire file content, giving us little control but
  // making it very simple. No need to close the file.
  content, err := ioutil.ReadFile("./input.txt")
  if err != nil {
      log.Fatal(err)
  }

  // Convert []byte to string
  text := string(content)
  // Move strings to array
  numbers := strings.Split(text, "\n")
  numberArray := make([]int, len(numbers))

  // Convert string array to int array
  for i := range numberArray {
    numberArray[i], _ = strconv.Atoi(numbers[i])
  }

  return numberArray
}

func fuelRequired(mass int) int {
  return (mass / 3) - 2
}

func partOne(masses []int) int {
  fuelNeeded := 0

  for _, n := range masses {
    fuelNeeded += fuelRequired(n)
  }

  return fuelNeeded
}

func partTwo(masses []int) int {
  fuelNeeded := 0

  for _, n := range masses {
    initialFuel := fuelRequired(n)
    currentFuel := fuelRequired(initialFuel)
    fuelNeeded += initialFuel

    for {
      if currentFuel <= 0 {
        break
      }

      fuelNeeded += currentFuel
      currentFuel = fuelRequired(currentFuel)
    }
  }

  return fuelNeeded
}

func main() {
  fmt.Println("The answer to part 1 is:", partOne(readMass()))
  fmt.Println("The answer to part 2 is:", partTwo(readMass()))
}
