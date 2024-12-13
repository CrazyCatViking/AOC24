package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
  leftNumber int
  rightNumber int
}

func main() {
  rules, updates := readInput("input.txt")

  parsedRules := parseRules(rules)
  parsedUpdates := parseUpdates(updates)

  part1(parsedUpdates, parsedRules)
}

func part1(updates [][]int, rules []Rule) {
  total := 0

  for _, update := range updates {
    isOk := true

    for _, rule := range rules {
      if !evalRule(rule, update) {
        isOk = false
        break
      }
    }

    if isOk {
      total += getMiddleValue(update)
    }
  }

  fmt.Println(total)
}

func getMiddleValue(update []int) int {
  middleIndex := len(update)/2
  return update[middleIndex]
}

func evalRule(rule Rule, update []int) bool {
  hasLeftNumber := slices.Contains(update, rule.leftNumber)
  hasRightNumber := slices.Contains(update, rule.rightNumber)

  if !hasLeftNumber || !hasRightNumber {
    return true
  }

  foundLeftNumber := false

  for _, value := range update {
    if (value == rule.leftNumber) {
      foundLeftNumber = true
    }

    if value == rule.rightNumber && !foundLeftNumber {
      return false
    }
  }

  return true
}

func parseUpdates(rawUpdates []string) [][]int {
  var parsedUpdates [][]int

  for _, rawUpdate := range rawUpdates {
    parts := strings.Split(rawUpdate, ",")

    var parsedUpdate []int

    for _, part := range parts {
      parsedPart, _ := strconv.Atoi(part)
      parsedUpdate = append(parsedUpdate, parsedPart)
    }

    parsedUpdates = append(parsedUpdates, parsedUpdate)
  }

  return parsedUpdates
}

func parseRules(rawRules []string) []Rule {
  var parsedRules []Rule

  for _, rawRule := range rawRules {
    parts := strings.Split(rawRule, "|")

    leftPage, _ := strconv.Atoi(parts[0])
    rightPage, _ := strconv.Atoi(parts[1])
    
    parsedRule := Rule{
      leftPage,
      rightPage,
    }

    parsedRules = append(parsedRules, parsedRule)
  }

  return parsedRules
}

func readInput(filePath string) ([]string, []string) {
  readFile, _ := os.Open(filePath)

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines) 

  var rules []string
  var updates []string

  isScanningRules := true

  for fileScanner.Scan() {
    line := fileScanner.Text()

    if (line == "") {
      isScanningRules = false
      continue
    }

    if (isScanningRules) {
      rules = append(rules, line)
    } else {
      updates = append(updates, line)
    }
  }

  readFile.Close()

  return rules, updates
}
