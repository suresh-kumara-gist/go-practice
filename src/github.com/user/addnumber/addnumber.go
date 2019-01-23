package main

import (
  "fmt"
)

var test, test2 bool



func main() {
  // Variables with initializers
  var i , j int = 1, 2

  fmt.Println(add(5, 6))
  var i int
  // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
  fmt.Println(addtype(5, 6))

  a,b := swap("world", "hello")
    fmt.Println(a,b)

   // Short variable declarations
  k := 3
  c, python, java := true, false, "no!"
}

func add(a int, b int ) int {
  return a + b
}

func addtype(a, b int) int {
  return a + b
}

func swap(x, y string) (string, string) {
    fmt.Println("I have returned multiple word")
  return  y , x
}


func split(sum int) (x, y int) {
  fmt.Println("// named rerurn values")
  x = sum * 4 / 9
  y = sum - x
  return
}
