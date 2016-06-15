package main

import "fmt"

func main() {
	x := []byte{'a', 'b', 'c'}
	x2p := &x[2]
	x = append(x, 'd', 'e', 'f')
	*x2p = 'Z'
	fmt.Println(x2p == &x[2], x2p, &x[2], string(*x2p), string(x))
}
