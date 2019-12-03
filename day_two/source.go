package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "strings"
  "strconv"
)

func readIntCode() []int {
  // Read entire file content, giving us little control but
  // making it very simple. No need to close the file.
  content, err := ioutil.ReadFile("./input.txt")
  if err != nil {
      log.Fatal(err)
  }

  // Convert []byte to string
  text := string(content)
  // Move strings to array
  numbers := strings.Split(text, ",")
  numberArray := make([]int, len(numbers))

  // Convert string array to int 2d array
  for i := range numberArray {
    numberArray[i], _ = strconv.Atoi(numbers[i])
  }

  return numberArray
}

func calculateProgram(n1 int, n2 int) []int {
  intCode := readIntCode()

  intCode[1] = n1
  intCode[2] = n2

  for n := range intCode {
    index := n * 4

    if n != 0 && index == 0 {
      continue
    }

    if intCode[index] == 1 {
      intCode[intCode[index + 3]] = intCode[intCode[index + 1]] + intCode[intCode[index + 2]]
    } else if intCode[index] == 2 {
      intCode[intCode[index + 3]] = intCode[intCode[index + 1]] * intCode[intCode[index + 2]]
    } else if intCode[index] == 99 {
      break
    }
  }

  return intCode
}

func partOne() int {
  return calculateProgram(12, 2)[0]
}

func partTwo() (int, int) {
  for x := 0; x <= 99; x++ {
    for y := 0; y <= 99; y++ {
      if calculateProgram(x, y)[0] == 19690720 {
        return x, y
      }
    }
  }

  return 0, 0
}

func main() {
  fmt.Println("The answer to part 1 is:", partOne())

  noun, verb :=  partTwo()
  fmt.Println("The answer to part 2 is:", noun, verb)
}
