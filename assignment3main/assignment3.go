package main

import (
	"fmt"
	"github.com/SalmanSaleh786/golang/golangexamples"
)

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	println(golangexamples.ConcatSlice(b))
	fmt.Printf("%s", golangexamples.Encrypt(b, 3))
	fmt.Printf("%s", golangexamples.EZGreetings("This is salman"))
}
