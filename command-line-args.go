Package main

import (
  “fmt”
  “OS”
)

func main () {
  fmt.Println(len(os.Args), os.Args)
  var a, b, c = 3, 4, "foo"
  fmt.Printf("a is of type %T\n", a)
}
