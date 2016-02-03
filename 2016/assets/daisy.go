// +build OMIT

package main

import "fmt"

func gopher(to, from chan int) {
	for {
		to <- 1 + <-from
	}
}

func main() {
	const n = 1000
	lastWhisperer := make(chan int)
	listener := lastWhisperer
	var whisperer chan int
	for i := 0; i < n; i++ {
		whisperer = make(chan int)
		go gopher(listener, whisperer)
		listener = whisperer
	}
	fmt.Printf("%d gopher() goroutines created\n", n)
	for i := 1; i < 10; i++ {
		whisperer <- i
		fmt.Println(<-lastWhisperer)
	}
}
