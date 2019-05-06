package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4} // unsorted
	sort.Float64s(a)
	fmt.Println(a)
	// Output: [1.1 2.2 3.3 4.4 5.5 6.6]
	a1 := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(a1)
	fmt.Println(a) // [1 2 3 4 5 6]
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s) // [6 5 4 3 2 1]

	sort.Ints(s)
	fmt.Println(s) // [6 5 4 3 2 1]
	a2 := []string{"PHP", "golang", "java", "python", "C", "Objective-C"}
	sort.Strings(a2)
	fmt.Println(a2) // [C Objective-C PHP golang java python]
}


