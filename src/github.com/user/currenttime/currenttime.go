package main

// this code groups the imports into a parenthesized, "factored" import statement.

import (
	"fmt"
	"time"
  "math"
)

func main() {
	fmt.Println("Current time is ", time.Now())
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

//  go install /Users/kumsu25/Sites/devdesktop/go-practice/src/github.com/user/currenttime/
// bin/currenttime


// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

// src/github.com/user/currenttime/currenttime.go:2:8: expected 'STRING', found '{'
// src/github.com/user/currenttime/currenttime.go:3:3: expected ';', found "fmt"
