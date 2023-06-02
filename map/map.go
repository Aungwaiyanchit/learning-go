package main

import "fmt"

func main() {
	var nilMap map[string]int     //map[keyType]valueType
	totalWins := map[string]int{} //we can declare map like that
	ages := make(map[int][]string, 10)

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	m := map[string]int{
		"hello": 4,
		"world": 0,
	}
	v, ok := m["hello"]

	intSet := map[int]bool{} //using map as sets
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(totalWins)
	fmt.Println(nilMap)
	fmt.Println("teams", teams["Orcas"])
	fmt.Println("ages of length", (ages))

	fmt.Println("zero", teams["3"])
	fmt.Println("v and ok", v, ok)
	v, ok = m["d"]
	fmt.Println("v and ok", v, ok)

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	//Summary
	/*
		Maps are like slices in several ways:
		• Maps automatically grow as you add key-value pairs to them.
		• If you know how many key-value pairs you plan to insert into a map, you can use
		make to create a map with a specific initial size.
		• Passing a map to the len function tells you the number of key-value pairs in a
		map.
		• The zero value for a map is nil.
		• Maps are not comparable. You can check if they are equal to nil, but you cannot
		check if two maps have identical keys and values using == or differ using !=.

	*/
}
