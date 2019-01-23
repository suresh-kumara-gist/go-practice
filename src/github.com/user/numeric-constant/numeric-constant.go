package main

import (
  "fmt"
)

func numberInt(x int) int{
   return x * 10 + 1
}

func numberFloat(x float64) float64{
   return x * 0.1
}

 const (
    BIG = 1 << 100
    SMALL = BIG >> 99
)

func main() {
  fmt.Println("***** Numeric constanst begin ******")

  fmt.Println(SMALL)

  fmt.Println(numberInt(SMALL))
  fmt.Println(numberFloat(SMALL))
  fmt.Println(numberFloat(BIG))

  fmt.Println()
  fmt.Println("***** Numeric constanst end ******")

  fmt.Println(check(BIG))
}

func check(BIG float64) float64{
  if BIG > 0 {
    return 1
  }
  return 0
}
