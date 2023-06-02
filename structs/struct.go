package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

type firstperson struct {
	name string
	age  int
}
type secondperson struct {
	name string
	age  int
}
type thirdPerson struct {
	age  int
	name string
}
type fourthPerson struct {
	firstName string
	age       int
}
type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}

func main() {
	adan := person{
		"Adan",
		20,
		"cat",
	}
	luci := person{
		"Luci",
		19,
		"dog",
	}
	fmt.Println(adan)
	fmt.Println(luci)
}
