package main

import (
  "fmt"
  "math"
)

func main() {
  // Count is a mapping from a perimeter to the number of integer-sided right triangles that can be formed
  count := make(map[int]int)

  for p := 0; p <= 1000; p++ {
    for i := 1; i < p; i++ {
      // Eliminate duplicate triangles by enforcing that i be the longest side
      for j := i; j < p - i; j++ {
        k := p - i - j
        // Eliminate duplicate triangles by enforcing that k be the shortest side
        if k < j {
          break
        }
        sq := i * i + j * j
        // Round to nearest integer
        sqrt := int(math.Sqrt(float64(sq)) + 0.5)
        // Check whether this is an integer square root
        if sqrt * sqrt != sq {
          continue
        }
        if k == sqrt {
          count[p]++
        }
      }
    }
  }

  // Creating a new mapping from number of triangles to perimeter
  revcount := make(map[int]int)
  for k, v := range count {
    revcount[v] = k
  }
  fmt.Println(revcount)
}
