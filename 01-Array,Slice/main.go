package main

import "fmt"

func main() {

	// define array
	var x = [5]int{1, 2, 3, 4, 5}
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(y)

	//define slice. slice is a dynamic array
	mySlice := []int{1, 2, 3, 4, 5}
	mySlice = append(mySlice, 6, 7, 8)
	fmt.Println(mySlice)

	// ---------------------------------------------------
	fmt.Println("-------------------------------------")

	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 10, 20, 30, 40, 50
	fmt.Println(s)

	s1 := s[1:3]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1[:cap(s1)]) // This id show we have memory leakage, thus // s1 := s[1:3] is not a good idea.

	// ----------------------------------------------

	fmt.Println("--------------------------")
	fmt.Println("Good implement for // s1 := s[1:3] without memory leakage.")
	s2 := make([]int, 2)
	copy(s2, s[1:3])
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s1[:cap(s2)])
}
