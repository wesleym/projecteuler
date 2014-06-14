package main

import (
  "fmt"
  "math"
)

func main() {
  count := make(map[int]int)
  for p := 0; p <= 1000; p++ {
    for i := 1; i < p; i++ {
      for j := i; j < p - i; j++ {
        k := p - i - j
        if k < j {
          break
        }
        sq := i * i + j * j
        sqrt := int(math.Sqrt(float64(sq)))
        if sqrt * sqrt != sq {
          continue
        }
        if k == sqrt {
          count[p]++
          //fmt.Println(i, j, k)
        }
      }
    }
  }
  revcount := make(map[int]int)
  for k, v := range count {
    revcount[v] = k
  }
  fmt.Println(revcount)
}
