package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
  readFile, _ := os.Open("input.txt")

  scanner := bufio.NewScanner(readFile)
  scanner.Split(bufio.ScanRunes)

  var runes []string

  for scanner.Scan() {
    runes = append(runes, scanner.Text())
  }

  part1(runes)
  part2(runes)
}

func part1(runes []string) {
  total := 0

  for len(runes) > 0 {
    if expect(&runes, "mul") {
      product, ok := tryParseMultiplication(&runes)

      if (ok) {
        total += product
      }
    }
  }

  fmt.Println(total)
}

func part2(runes []string) {
  total := 0
  previousConditional := "do"

  for len(runes) > 0 {
    token, ok := parse(&runes)

    if !ok {
      continue;
    }

    if (token != "mul") {
      previousConditional = token
    }

    if (previousConditional == "don't") {
      continue;
    }

    product, ok := tryParseMultiplication(&runes)

    if (ok) {
      total += product
    }
  }

  fmt.Println(total)
}

func tryParseMultiplication(runes *[]string) (int, bool) {
  if !expect(runes, "(") {
    return 0, false
  }

  leftValue := getValue(runes, ",")
  leftNum, err := strconv.Atoi(leftValue)

  if err != nil {
    return 0, false
  }

  if !expect(runes, ",") {
    return 0, false
  }

  rightValue := getValue(runes, ")")
  rightNum, err := strconv.Atoi(rightValue)

  if err != nil {
    return 0, false
  }

  if !expect(runes, ")") {
    return 0, false
  }

  product := leftNum * rightNum
  return product, true
}

func parse(runes *[]string) (string, bool) {
  tokens := []string{"don't", "do", "mul"}

  for _, token := range tokens {
    parsedToken := strings.Join((*runes)[0 : len(token)], "")

    if (parsedToken == token) {
      *runes = (*runes)[len(token):]
      return parsedToken, true
    }
  }

  currentRune := (*runes)[0]
  *runes = (*runes)[1:]

  return currentRune, false
}

func expect(runes *[]string, token string) bool {
  parsedToken := strings.Join((*runes)[0 : len(token)], "")

  if (parsedToken == token) {
    *runes = (*runes)[len(token):]
    return true;
  }

  *runes = (*runes)[1:]
  return false
}

func getValue(runes *[]string, delimiter string) string {
  var value []string

  for _, r := range *runes {
    if !unicode.IsDigit(rune(r[0])) {
      *runes = (*runes)[len(value):]
      break;
    }

    if r == delimiter {
      *runes = (*runes)[len(value):]
      break;
    }

    value = append(value, r)
  }

  return strings.Join(value, "")
}
