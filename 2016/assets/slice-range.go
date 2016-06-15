package main

import "fmt"

func main() {
	slice := []int{11, 22, 33, 44, 55}
	for key, value := range slice {
		fmt.Println(key, value, slice[key])
		value >>= 1
		fmt.Println(">>", value, slice[key])
	}
}
