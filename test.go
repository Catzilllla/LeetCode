package main

import (
	"fmt"
)

func main() {
	var x []int
	var y = []int{}
	makeSlice := make([]int, 2)
	fmt.Println(&x)
	x = append(x, 5, 4, 3, 2, 1)
	fmt.Println(&x)
	x = append(x, 5, 4, 3, 2, 1)
	fmt.Println(&x)
	slice := []int{10, 20, 30}
	fmt.Println(&slice)

	y = append(y, 1, 3, 2)
	_ = copy(makeSlice, y)
	makeSlice = append(makeSlice, x[1:]...)
	fmt.Println("SLICE\n")
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println("\n")
	fmt.Println("x: ", cap(x), len(x), x, x[0])
	fmt.Println("y: ", cap(y), len(y), y, y[0])
	fmt.Println("makeSlice: ", cap(makeSlice), len(makeSlice), makeSlice)
	fmt.Printf("\nMAPS\n")

	var nilMap map[string]int
	m := map[string]int {
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println("v, ok: ", v, ok)
	v1, ok1 := m[""]
	fmt.Println("v, ok: ", v1, ok1)

	m["new"] = 100
	v2, ok2 := m["new"]
	delete(m, "new")
	fmt.Println("v, ok: ", v2, ok2, m)
	clear(m)
	fmt.Println(m)

	// PANIC
	// nilMap["1"] = 1
	fmt.Println(nilMap == nil)
	fmt.Println(nilMap, nilMap[""], len(nilMap))
	total := map[string]int{}
	fmt.Println(total == nil)
	total["ABC"] = 1
	fmt.Println(total == nil)
	fmt.Println(total, total["F"], total["ABC"], len(total))

	totalInt := map[int]int{}
	fmt.Println(totalInt == nil)
	totalInt[100] = 1
	fmt.Println(totalInt == nil)
	fmt.Println(totalInt, totalInt[0], totalInt[100], len(totalInt))

	xArrayPointer := (*[4]int)(x)
	fmt.Println("xArrayPointer: ", xArrayPointer)

	fmt.Println("\n||||||||||||||||||||||||\n")
	fmt.Println("SET on MAP\n")

	intSet := map[int]bool{}
	intSetStruct := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
	fmt.Println(len(vals), len(intSet))

	for _, v := range vals {
		intSetStruct[v] = struct{}{}
	}
	n := 5
	fmt.Println(intSetStruct)
	if _, ok := intSetStruct[n]; ok {
		fmt.Printf("%d is in the SET", n)
	}

	fmt.Println("\n||||||||||||||||||||||||\n")
	fmt.Println("\nFOR МЕТКИ")

	sample := []string{"hello", "apple_dsfds"}
	outer:
		for _, sample := range sample {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
				if r == 'l' {
					continue outer
				}
			}
			fmt.Println()
		}
}
