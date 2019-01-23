package main

import (
  "fmt"
)

type Vertex struct {
    X int
    Y int
}


var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)



func main() {
  i := 3
  p := &i
  fmt.Println(p)
  fmt.Println(*p)

  fmt.Println(Vertex{1, 2})
  v := Vertex{3, 4}
  fmt.Println(v.X)
  fmt.Println(v.Y)

  // pointers to struct

  l := Vertex{8,7}
  m := &l

  fmt.Println(m.X)
  fmt.Println(m.Y)

  prime := [6]int{12,13,14, 15, 16, 17}

  var s []int = prime[1:4]
  fmt.Println(s)


  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)

  a := names[0:2]
  b := names[1:4]
  fmt.Println(a, b)

  b[0] = "XXXXX"
  fmt.Println(a, b)
  fmt.Println(names)


  q := []int{1,2,3,4,5,6}
  fmt.Println(q)

  ss := []struct {
  		i int
  		b bool
  	}{
  		{2, true},
  		{3, false},
  		{5, true},
  		{7, true},
  		{11, false},
  		{13, true},
  	}

  fmt.Println(ss)


  wq := []int{2, 3, 5, 7, 11, 13}
  	printSlice(wq)

  	// Slice the slice to give it zero length.
  	wq = wq[:0]
  	printSlice(wq)

  	// Extend its length.
  	wq = wq[:4]
  	printSlice(wq)

  	// Drop its first two values.
  	wq = wq[2:]
  	printSlice(wq)

    var niltest []int
    fmt.Println(niltest, len(niltest), cap(niltest))
    if niltest == nil {
      fmt.Println("nil!")
    }
    // Nil slices


    ty := make([]int, 5)
    printSlice1("a", ty)

    po := make([]int, 0, 5)
    printSlice1("b", po)

    cr := po[:2]
    printSlice1("c", cr)

    de := cr[2:5]
    printSlice1("d", de)

}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
