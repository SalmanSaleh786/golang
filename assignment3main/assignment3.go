package main

import (
	"fmt"
	"github.com/SalmanSaleh786/golangexamples"
)

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	println(golangexamples.ConcatSlice(b))
	fmt.Printf("%d", golangexamples.Encrypt(b, 3))
	fmt.Printf("%s", golangexamples.EZGreetings("This is salman"))
}
