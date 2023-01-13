package main

import "fmt"

func main() {

	// Pointer ----------------------------------------------------------------
	fmt.Println("Pointer -------------------------------------------")

	var pI *int //a memory address with type of int
	var I = 3
	// *pI -> this is value of pI
	// pI -> this is address
	// I -> this is a value
	// &I -> this is address of I

	pI = &I
	fmt.Println(pI)
	fmt.Println(*pI)
	inc(pI)   // send pI itself
	inc0(*pI) // send a copy of pI
	fmt.Println(*pI)

	// Defer ----------------------------------------------------------------
	fmt.Println("Defer ------------------------------------------")
	testDefer()

	// Panic ----------------------------------------------------------------
	fmt.Println("Panic -------------------------------------------")

	fmt.Println("Before panic")
	// testPanic()
	handlePanic()
	fmt.Println("After panic")

}

func inc(v *int) {
	*v++
}

func inc0(v int) {
	v++
}

func testDefer() {
	// After running normal codes of a function, defers will be running
	defer fmt.Println("test 1")
	defer fmt.Println("test 2")
	fmt.Println("test 3")
	fmt.Println("test 4")
}

func testPanic() {
	// use panic() to throwing an exception.
	panic("I am panic!!!") // Program will be crashed at this line.
}

func handlePanic() {
	defer func() {
		if err := recover(); err != nil {
			// recover() getting the error.
			fmt.Println("Panic handled : ")
			fmt.Println(err)
		}
	}()                    // func(){}() -> Execute func at define line.
	panic("I am panic!!!") // Program not crash at this line because we used defer like try-catch.

}
