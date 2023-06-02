package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(10)
	for a < 10 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}

func shadow_variable() {
	//shadowing variable
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func shadowing_mutlivalue() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

func forFourWays() {

	//first way like C-style
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	//seocond way "conditon-only-for statement"
	i := 1
	for i < 2 {
		fmt.Println(i)
	}

	// "for-range"
	oddVals := []int{3, 7, 5, 5, 3, 3}
	for i, v := range oddVals {
		fmt.Println(i, v)
	}

}
