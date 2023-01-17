package main

import "fmt"

type TestInter interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
	IncrementP()
	GetInternalValueP() int
}

// NewTestInter Constructor
func NewTestInter() TestInter {
	return new(tStruct)
}

// NewTestInterWithV Constructor
func NewTestInterWithV(v int) TestInter {
	return &tStruct{i: v}
}

type tStruct struct {
	i int
}

func (ts tStruct) SayHello() {
	fmt.Println("Hello!")
}
func (ts tStruct) Say(s string) {
	fmt.Println(s)
}

func (ts tStruct) Increment() {
	ts.i++
}

func (ts tStruct) GetInternalValue() int {
	return ts.i
}

func (ts *tStruct) IncrementP() {
	ts.i++
}

func (ts *tStruct) GetInternalValueP() int {
	return ts.i
}

type EmbeddingStruct struct {
	*tStruct
}

// Empty interface is parent of all objects in golang
func testEmpty(v interface{}) {
	//fmt.Println(v)

	//if v, ok := v.(int); ok { // This is type assertion
	//	fmt.Println("I am int and my value is ", v)
	//} else {
	//	fmt.Println("I am not int, my value is ", v)
	//}

	switch val := v.(type) { // This is type switch
	case int:
		fmt.Println("I am int and my value is ", val)
	case string:
		fmt.Println("I am string and my value is ", val)
	default:
		fmt.Println("I am not string and int, my value is ", val)
	}
}

func main() {
	var test TestInter
	//test = &tStruct{} // new(tStruct)
	//test = NewTestInter() // new(tStruct)
	test = NewTestInterWithV(5) // new(tStruct)
	test.SayHello()
	test.Say("Hello my friend!")

	// Without pointer (Sends a copy of test to methods)
	test.Increment()
	test.Increment()
	test.Increment()
	fmt.Println(test.GetInternalValue())

	// With pointer
	test.IncrementP()
	test.IncrementP()
	test.IncrementP()
	fmt.Println(test.GetInternalValue())

	// Embedding interface -> Access to interface methods from another method
	fmt.Println("Embedding interface --------------")
	te := EmbeddingStruct{&tStruct{i: 80}}
	te.Say("Embedding hahahaaaa!")
	te.IncrementP()
	te.IncrementP()
	te.IncrementP()
	te.IncrementP()
	fmt.Println(te.GetInternalValueP())

	// Empty interface
	fmt.Println("Empty interface --------------")
	testEmpty(5)
	testEmpty("This is a test for empty interface.")
	testEmpty(&test)

}
