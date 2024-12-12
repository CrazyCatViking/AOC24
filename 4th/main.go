package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  lines := readInput("input.txt")
  part1(lines)
  part2(lines)
}

func part2(lines []string) {
  total := 0

  m := len(lines)
  n := len(lines[0])

  for x := 0; x < m - 2; x++ {
    for y := 0; y < n - 2; y++ {
      sector := make([]string, 3)
      sector[0] = lines[x][y:y+3]
      sector[1] = lines[x+1][y:y+3]
      sector[2] = lines[x+2][y:y+3]

      if (checkSector(sector)) {
        total++
      }
    }
  }

  fmt.Println(total)
}

func checkSector(sector []string) bool {
  diag1 := ""
  diag2 := ""

  for x := 0; x < 3; x++ {
    diag1 = diag1 + string(sector[x][x])
    diag2 = diag2 + string(sector[x][2 - x])
  }

  return isMas(diag1) && isMas(diag2)
}

func isMas(token string) bool {
  return token == "MAS" || token == "SAM"
}

func part1(lines []string) {
  total := 0

  total += checkHorizontalWords(lines)
  total += checkVerticalWords(lines)
  total += checkDiagonalWords(lines)
  total += checkDiagonalWords2(lines)

  fmt.Println(total)
}

func checkHorizontalWords(lines []string) int {
  total := 0

  for i := 0; i < len(lines); i++ {
    token := ""

    for j := 0; j < len(lines[i]); j++ {
      cur := lines[i][j];

      token = token + string(cur)

      if (len(token) == 4) {
        j = j - 3

        if (token == "XMAS" || token == "SAMX") {
          total++
        }

        token = ""
      }
    }  
  }

  fmt.Println(total)

  return total
}

func checkVerticalWords(lines []string) int {
  total := 0

  m := len(lines)
  n := len(lines[0])

  i := 0
  j := 0

  token := ""

  for i < m && j < n {
    token = token + string(lines[i][j])

    if len(token) == 4 {
      i = i - 3

      if (token == "XMAS" || token == "SAMX") {
        total++
      }

      token = ""
    }

    if i == m - 1 {
      i = 0
      j++

      token = ""

      continue
    }

    i++
  }

  fmt.Println(total)

  return total
}

func checkDiagonalWords(lines []string) int {
  total := 0

  m := len(lines)
  n := len(lines[0])

  for d := range n+m-1 {
    token := ""

    for x := max(0, d-m+1); x < min(n, d+1); x++ {
      cur := string(lines[x][d - x])
      token = cur + token

      if len(token) == 4 {
        if token == "XMAS" || token == "SAMX" {
          total++
        }

        token = ""
        x -= 3
      }
    }
  }

  fmt.Println(total)

  return total
}

func checkDiagonalWords2(lines []string) int {
  total := 0

  m := len(lines)
  n := len(lines[0])

  for d := range n+m-1 {
    token := ""

    for x := max(0, d-m+1); x < min(n, d+1); x++ {
      cur := string(lines[x][n - 1 - d + x])

      token = cur + token

      if len(token) == 4 {
        if token == "XMAS" || token == "SAMX" {
          total++
        }

        token = ""
        x -= 3
      }
    }
  }

  fmt.Println(total)

  return total
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
