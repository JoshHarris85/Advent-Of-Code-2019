package main

import (
  "fmt"
  "strconv"
)

func findMatches(current, max int, countTriple bool) []int {
  var found []int

  for x := current; x <= max; x++ {
    currentString := strconv.Itoa(x)
    incrementOrdered := true
    hasDouble := false

    for y := 1; y < len(currentString); y++ {
      var nextNum, lastCompareNum int
      currentNum, _ := strconv.Atoi(currentString[y:(y+1)])
      previousNum, _ := strconv.Atoi(currentString[(y-1):y])

      if countTriple {
        if y < len(currentString) - 1 {
          nextNum, _ = strconv.Atoi(currentString[(y+1):(y+2)])
        } else {
          nextNum, _ = strconv.Atoi(currentString[y+1:])
        }

        if y != 1 {
          lastCompareNum, _ = strconv.Atoi(currentString[(y-2):(y-1)])
        } else {
          lastCompareNum = -99
        }
      }

      if !hasDouble && (currentNum == previousNum && !countTriple) || (!hasDouble && currentNum == previousNum && currentNum != nextNum && previousNum != lastCompareNum) {
        hasDouble = true
      }

      if incrementOrdered && previousNum > currentNum {
        incrementOrdered = false
      }
    }

    if incrementOrdered && hasDouble {
      found = append(found, x)
    }
  }

  return found
}

func partOne() int {
  return len(findMatches(307237, 769058, false))
}

func partTwo() int {
  return len(findMatches(307237, 769058, true))
}

func main() {
  fmt.Println("The answer to part 1 is:", partOne())
  fmt.Println("The answer to part 2 is:", partTwo())
}
