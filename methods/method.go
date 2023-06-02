package main

import (
	"fmt"
	"time"
)


func main() {
	demo()
}


type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

//pointers receivers and value receivers

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { //don't do this
	c.Increment()
	fmt.Println("in doupdateWrong", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doupdateRight", c.String())
}

//


type IntTree struct {
	val int
	left, right * IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val > it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Conatins(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Conatins(val)
	case val > it.val:
		return it.right.Conatins(val)
	default:
		return true
	}
}