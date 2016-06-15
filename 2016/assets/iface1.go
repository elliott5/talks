package main

import "strconv"
import "fmt"

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

type Stringer interface {
	String() string
}

func main() {
	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s)
}
