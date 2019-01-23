package main

import (
  "fmt"
  "math/cmplx"
)

var (

  ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  // formatted outuput
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


  fmt.Println("////////")
  fmt.Println("zero values")
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias for uint8
//
// rune // alias for int32
//      // represents a Unicode code point
//
// float32 float64
//
// complex64 complex128
