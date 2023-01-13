package main

import (
	"fmt"
)

func compare(x int, y int) (isEqual bool, difference int) {
	//if x > y {
	//	isEqual = false
	//	difference = x - y
	//} else if y > x {
	//	isEqual = false
	//	difference = y - x
	//} else {
	//	isEqual = true
	//	difference = 0
	//}

	switch {
	case x > y:
		isEqual = false
		difference = x - y
	case y > x:
		isEqual = false
		difference = y - x
	default:
		isEqual = true
		difference = 0
	}

	return
}

//define parameter globaly
var c int = 6

func main() {
	fmt.Printf("Hello world.\n")
	fmt.Println(compare(12, 45))
	fmt.Println(compare(74, 45))
	fmt.Println(compare(45, 45))

	//x := 5

	switch c {
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
		fallthrough //this means that run next case too.
	default:
		fmt.Println("default")
	}

	for x := 0; x <= 5; x++ {
		// do nothing
	}

	switch x := 0; x {
	
	}

}
