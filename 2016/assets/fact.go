package main // example inspired by gobyexample.com/recursion

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // recursion: see recursion
}

func main() {
	for f := 1; f <= 10; f++ {
		fmt.Println(fact(f))
	}
}
