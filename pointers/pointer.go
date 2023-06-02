package main

import (
	"fmt"
)


func main() {
	// var f *int
	x := []int {1, 2, 4}
	a := x
	fmt.Println(x)
	fmt.Println(a)
	x[1] = 23
	a = append(a, 3, 4, 5, 5)
	fmt.Println(x)
	fmt.Println(a)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Println(g)
}
