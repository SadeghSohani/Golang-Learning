package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bad implement -------------------------------")
	goRoutine()
	// ------------------------------------
	fmt.Println("Good implement -------------------------------")
	testChannel()
	// ------------------------------------
	fmt.Println("Send signal from go routine -------------------------------")
	ch := make(chan int)
	go testSignal(ch)
	for i := range ch {
		fmt.Println(i)
	}

	// Buffer *****************************************
	fmt.Println("Buffer ******************************")
	buffer()
	// ---------------
	fmt.Println("-------------")
	select {
	case v1 := <-WaitChannel(5, 2):
		{
			fmt.Println(v1)
		}
	case v2 := <-WaitChannel(6, 2):
		{
			fmt.Println(v2)
		}
		//default:
		//	{
		//		fmt.Println("Dont wait for go routines (v1, v2).")
		//	}
	}

	// ---------------
	fmt.Println("-------------")

	tick := time.Tick(100 * time.Millisecond)  // tick is a channel
	boom := time.After(500 * time.Millisecond) // boom is a channel
	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("Boom!!")
			return
		default:
			fmt.Println("*|*")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func goRoutine() {

	go fmt.Println("Hello") //this is in another thread.
	fmt.Println("World")
	time.Sleep(3 * time.Second)

}

// ------------------------------------
func testChannel() {
	c := make(chan bool) // c is a channel that return boolean
	go func() {
		fmt.Println("Hello")
		c <- true
	}()
	<-c // this line stop the program until go routine finish.
	fmt.Println("World")
}

func testSignal(c chan int) {
	i := 0
	for i <= 5 {
		c <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(c) // This is important for finish process in main thread
}

// ********************************
func buffer() {
	buffer := make(chan string, 2)
	buffer <- "Hello"
	buffer <- "world"
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)

}

func WaitChannel(v, i int) chan int {
	channel := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		channel <- v
	}()
	return channel
}
