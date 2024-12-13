package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector struct {
  x int
  y int
}

var up = Vector{ x: 0, y: -1 }
var down = Vector{ x: 0, y: 1 }
var left = Vector{ x: -1, y: 0 }
var right = Vector{ x: 1, y: 0 }

func main() {
  part1()
  part2()
}

func part1() {
  mapMatrix := getMap("input.txt")

  initialPos, initialDir := findInitialPos(mapMatrix)
  mapMatrix[initialPos.y][initialPos.x] = '.'

  currentPos := initialPos
  currentDir := initialDir

  visitedPos := make(map[Vector]bool)

  for true {
    nextPos, nextDir, validPos := moveNext(mapMatrix, currentPos, currentDir)
    
    if !validPos {
      break;
    }

    currentPos = nextPos
    currentDir = nextDir

    visitedPos[currentPos] = true
  }

  fmt.Println(len(visitedPos))
}

func part2() {
  mapMatrix := getMap("input.txt")

  initialPos, initialDir := findInitialPos(mapMatrix)
  mapMatrix[initialPos.y][initialPos.x] = '.'

  numLoops := 0

  // This is an unga bunga brute force way to solve it xD
  // Computer go brrRRRrrRrrRRRRRrr
  for i := 0; i < len(mapMatrix); i++ {
    for j := 0; j < len(mapMatrix[i]); j++ {
      if initialPos.x == j && initialPos.y == i {
        continue
      } 

      if mapMatrix[i][j] == '#' {
        continue
      }
  
      var matrixCopy = copyMatrix(mapMatrix) 
      matrixCopy[i][j] = '#'

      if isInLoop(matrixCopy, initialPos, initialDir) {
        numLoops++
      }
    }
  }

  fmt.Println(numLoops)
}

func isInLoop(mapMatrix [][]rune, initialPos Vector, initialDir Vector) bool {
  currentPos := initialPos
  currentDir := initialDir

  visitedPos := make(map[Vector]Vector)

  for true {
    nextPos, nextDir, validPos := moveNext(mapMatrix, currentPos, currentDir)
    
    if !validPos {
      break;
    }

    currentPos = nextPos
    currentDir = nextDir

    direction, hasSeenPos := visitedPos[currentPos]

    if !hasSeenPos {
      visitedPos[currentPos] = currentDir
      continue
    }

    if currentDir.x == direction.x && currentDir.y == direction.y {
      return true
    }
  }

  return false
}

func moveNext(mapMatrix [][]rune, currentPos Vector, currentDir Vector) (Vector, Vector, bool) {
  nextPos := Vector{ x: currentPos.x + currentDir.x, y: currentPos.y + currentDir.y }  

  if nextPos.x > len(mapMatrix) - 1 || nextPos.x < 0 || nextPos.y > len(mapMatrix[0]) - 1 || nextPos.y < 0 {
    return nextPos, currentDir, false
  }

  nextPosRune := mapMatrix[nextPos.y][nextPos.x]

  switch nextPosRune {
  case '.':
    return nextPos, currentDir, true
  case '#':
    nextDir := getNextDir(currentDir)
    return moveNext(mapMatrix, currentPos, nextDir)
  default:
    return nextPos, currentDir, false
  }
}

func getNextDir(currentDir Vector) Vector {
  switch currentDir {
  case up:
    return right
  case right:
    return down
  case down:
    return left
  case right:
    return up
  default:
    return up
  }
}

func findInitialPos(mapMatrix [][]rune) (Vector, Vector) {
  for i := 0; i < len(mapMatrix); i++ {
    for j := 0; j < len(mapMatrix[i]); j++ {
      if mapMatrix[i][j] == '^' {
        return Vector{ x: j, y: i }, up
      }
    }
  }

  return Vector{ x: -1, y: -1 }, up
}

func copyMatrix(matrix [][]rune) [][]rune {
  matrixCopy := make([][]rune, len(matrix))
  for i := range matrix {
    matrixCopy[i] = make([]rune, len(matrix[i]))
    copy(matrixCopy[i], matrix[i])
  }

  return matrixCopy
}

func getMap(path string) [][]rune {
  lines := readInput(path)

  mapMatrix := make([][]rune, len(lines))

  for i, line := range lines {
    mapMatrix[i] = make([]rune, len(line))

    for j, char := range line {
      mapMatrix[i][j] = char
    }
  }

  return mapMatrix
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

