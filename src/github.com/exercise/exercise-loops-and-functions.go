package main

import (
  "fmt"
)

func Sqrt(x float64) float64{
  //  we want to find the number z for which z² is most nearly x.
  // for z:=0;  z++ {
  //   if z^2 == x {
  //     fmt.Println(z)
  //     break
  //   }
  // }

  z:=float64(1.0);
  for i:=0; i< 10;i++  {
    v := z
    z -= (z*z - x) / (2*z)
    fmt.Println(z)
    if z == v {
      fmt.Println("I found value")
      break
    }
  }
  // fmt.Println(z)
  return z
}

func main() {
	fmt.Println(Sqrt(4))
}
