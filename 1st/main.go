package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
  taskA()
  taskB()
}

func taskA() {
  sort.Ints(ListA)
  sort.Ints(ListB)

  totalDelta := 0 

  for i := 0; i < len(ListA); i++ {
    totalDelta += int(math.Abs(float64(ListA[i] - ListB[i])))
  }
  
  fmt.Println("Total Delta:", totalDelta)
}

func taskB() {
  total := 0

  for _, valA := range ListA {
    count := 0

    for _, valB := range ListB {
      if valA == valB {
        count++
      }
    }

    total += valA * count
  }

  fmt.Println("Total:", total)
}

