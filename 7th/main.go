package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
  Division Operation =  iota
  Subtraction Operation = iota
)

type Equation struct {
  sum int
  numbers []int
}

func main() {
  part1()
  part2()
}

func part1() {
  input := readInput("input.txt")
  equations := parseInput(input)

  total := 0

  for _, equation := range equations {

    if testEquation(equation.sum, 0, equation.numbers) {
      total += equation.sum
    }
  }

  fmt.Println(total)
}

func part2() {
  input := readInput("input.txt")
  equations := parseInput(input)

  total := 0

  for _, equation := range equations {

    if testEquation2(equation.sum, 0, equation.numbers) {
      total += equation.sum
    }
  }

  fmt.Println(total)
}

func testEquation(targetSum int, currentSum int, numbers []int) bool {
  if len(numbers) == 0 && targetSum == currentSum {
    return true
  }

  if len(numbers) == 0 && targetSum != currentSum {
    return false
  }

  multiplyRes := testEquation(targetSum, currentSum * numbers[0], numbers[1:])
  divisionRes := testEquation(targetSum, currentSum + numbers[0], numbers[1:])

  return multiplyRes || divisionRes
}

func testEquation2(targetSum int, currentSum int, numbers []int) bool {
  if len(numbers) == 0 && targetSum == currentSum {
    return true
  }

  if len(numbers) == 0 && targetSum != currentSum {
    return false
  }

  multiplyRes := testEquation2(targetSum, currentSum * numbers[0], numbers[1:])
  divisionRes := testEquation2(targetSum, currentSum + numbers[0], numbers[1:])
  concatRes := testEquation2(targetSum, concatNumbers(currentSum, numbers[0]), numbers[1:])

  return multiplyRes || divisionRes || concatRes
}

func concatNumbers(left int, right int) int {
  leftString := strconv.Itoa(left)
  rightString := strconv.Itoa(right)

  concatNumber, _ := strconv.Atoi(leftString + rightString)

  return concatNumber
}

func parseInput(lines []string) []Equation {
  equations := make([]Equation, len(lines)) 

  for i, line := range lines {
    parts := strings.Split(line, ":")

    sum, _ := strconv.Atoi(parts[0])

    numberString := strings.Trim(parts[1], " ") 

    var numbers []int

    for _, num := range strings.Split(numberString, " ") {
      number, _ := strconv.Atoi(string(num))
      numbers = append(numbers, number) 
    }

    equations[i] = Equation{ sum, numbers }
  }

  return equations
}

func readInput(filePath string) []string {
  readFile, _ := os.Open(filePath)

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines) 

  var lines []string

  for fileScanner.Scan() {
    line := fileScanner.Text()
    lines = append(lines, line)
  }

  readFile.Close()

  return lines
}

