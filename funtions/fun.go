package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	result, remainder, err := divAndRemainder4(19, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	fmt.Println(multiple(4, 9))
	fmt.Println(result, remainder, err)
	fmt.Println(opMap)

	test()
	DeferDemo()

	p := person{}
	i := 2
	s := "Hello"

	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	//Anonymous Functions
	func(j int) {
		fmt.Println(j)
	}(2)
}

// using a struct to simulate named parameters
func MyFunc(opts MyFuncOpts) error {
	return fmt.Errorf("d")
}

//

// delcare function in go is  keyword "func"
func div(numberator int, denominator int) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numberator / denominator)
}

//

// basic func
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

//

// return mutiples vlaues
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot diveide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

//

// you can also named the return vals like outher languages
func divAndRemainder2(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divided by zero")
		os.Exit(1)
	}
	result, remainder, err = numerator/denominator, numerator%denominator, nil
	return result, remainder, err
}

//

// named the return vals are good but it has some bad cases, like shadowing variables
func divAndRemainder3(numerator int, denomerator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denomerator == 0 {
		err = errors.New("cannot divided by zero")
		return 0, 0, err
	}
	return numerator / denomerator, numerator % denomerator, nil
}

//

// blank_reutrn (something called naked)     //notice -> shouldn't use
func divAndRemainder4(numerator int, denomerator int) (result int, remainder int, err error) {
	if denomerator == 0 {
		err = errors.New("cannot divided by zero")
		return
	}
	result, remainder = numerator/denomerator, numerator%denomerator
	return
}

//

// just like other languages funcitons are values
func multiple(i int, j int) int { return i * j }

//

// function type declationg
type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"add": func(i, j int) int { return i + j },
	"sub": func(i, j int) int { return i - j },
	"div": func(i, j int) int { return i / j },
}

//

// closure (function declared inside of funtions)
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func test() {

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].FirstName
	})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}

//

//return func from func

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// defer
func DeferDemo() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

}

// go is call by value
type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}
