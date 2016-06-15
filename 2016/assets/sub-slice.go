package main

import "fmt"

func main() {
	x := []int{2, 3, 5, 7, 11}
	y := x[1:3]
	fmt.Println(x, y)
	y[1] = 55
	fmt.Println(x, y)
	y = append(y, 77)
	fmt.Println(x, y)
}
