package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ---- interface ByName
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// ---- interface ByAge
type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
		{"Jim", 14},
		{"Jess", 15},
		{"Jemima", 11},
	}

	sort.Sort(ByName(kids))
	// sort.Sort(kids)
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
