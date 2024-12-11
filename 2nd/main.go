package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  part1()
  part2()
}

func part1() {
  lines := readInput("input.txt")

  numSafeReports := 0

  for _, line := range lines {
    report := strings.Split(line, " ")

    if (evaluateLine(report)) {
      numSafeReports++
    }
  }

  fmt.Println(numSafeReports)
}

func part2() {
  lines := readInput("input.txt")

  numSafeReports := 0

  for _, line := range lines {
    report := strings.Split(line, " ")

    if (evaluateLine(report)) {
      numSafeReports++
      continue
    } else {
      for i := range report {
        simplifiedReport := make([]string, 0, len(report)-1)
        simplifiedReport = append(simplifiedReport, report[:i]...)
        simplifiedReport = append(simplifiedReport, report[i+1:]...)

        if (evaluateLine(simplifiedReport)) {
          numSafeReports++
          break;
        }
      }
    }
  }

  fmt.Println(numSafeReports)
}

func evaluateLine(report []string) bool {
  isSafe := true;
  direction := 0;

  for index, value := range report {
    if (index + 1 == len(report)) {
      continue
    }

    currVal, _ := strconv.Atoi(value)
    nextVal, _ := strconv.Atoi(report[index + 1])

    diff := currVal - nextVal
    currDirection := 0

    if (diff > 0) {
      currDirection = 1
    } else if (diff < 0) {
      currDirection = -1
    }

    if direction != 0 && currDirection != direction {
      isSafe = false
      break;
    }

    absDiff := abs(diff)

    if (absDiff > 3 || absDiff < 1) {
      isSafe = false
      break;
    }

    direction = currDirection
  }

  return isSafe
}

func isSumSafe(sum int, dir int) bool {
  absSum := abs(sum)
  return sum != 0 && absSum <= 3 && absSum >= 1 && sum / absSum == dir 
}

func abs(x int) int {
  if (x > 0) {
    return x
  }

  return -x
}

func readInput(filePath string) []string {
  readFile, _ := os.Open(filePath)

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines) 

  var lines []string

  for fileScanner.Scan() {
    lines = append(lines, fileScanner.Text())
  }

  readFile.Close()

  return lines
}
