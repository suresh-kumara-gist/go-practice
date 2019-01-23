package main

import (
  "fmt"
  "math"
)

type Myfloat float64

type Vertex struct {
  X, Y  float64
}

func (v Vertex)  Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (t Myfloat)  Abs2() float64 {
  return math.Sqrt(float64(t))
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func (v *Vertex)  Abs4(f float64) float64 {
  v.X := v.X * f
  v.Y := v.Y * f
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  fmt.Println(Abs(v))


  t := Myfloat(64)
  fmt.Println(t.Abs2())


  s := Vertex{7, 8}
  fmt.Println(s.Abs4(10))
  fmt.Println(s)

}

// bin/methods
//  go install src/github.com/advanced/methods.go
