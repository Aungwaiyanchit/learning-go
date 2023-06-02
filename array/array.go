package main

import "fmt"

func main() {
//  array()
var s string = "Hello 😒"
var bs []byte  = []byte(s)
var rs []rune = []rune(s)

fmt.Println(len(s))
fmt.Println(bs)
fmt.Println(rs)
}



func array() {
	var y []int
	var x = [3]int{1, 2, 3} // array "need to assign the size of array"
	var c = []int{1, 2, 3} //slice "don't need to assign the size of the slice"
	var z = []int{4, 5, 6}
	b := []string{"adan", "lucu", "jenifer"}
	a := make([]int, 4)
	a = append(a, 10)
	arr1 := []int{1, 2, 3, 4}
	arr2 := make([]int, 4)
	num := copy(arr2, arr1)


	fmt.Println("num & arr2", num, arr2)

	fmt.Println("before" , b)
	b = b[:2]
	println("a", a)

	fmt.Println("after slice", b)
	fmt.Println(y)
	y = append(y, 1)
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println((len(x))) 
	fmt.Println(y)
	fmt.Println(append(c, z...)) //append the slices 
	fmt.Println(cap(z))
}
