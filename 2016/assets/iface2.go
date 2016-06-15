package main

import "strconv"
import "fmt"
import "time"

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

type Stringer interface {
	String() string
}

func main() {
	b, s, itters := Binary(200), Stringer(Binary(200)), 10000000

	l, n := itters, time.Now()
	for ; l > 0; l-- {
		_ = strconv.FormatUint(uint64(b), 2)
	}
	fmt.Println("no method call", time.Since(n))

	l, n = itters, time.Now()
	for ; l > 0; l-- {
		_ = b.String()
	}
	fmt.Println("direct method call", time.Since(n))

	l, n = itters, time.Now()
	for ; l > 0; l-- {
		_ = s.String()
	}
	fmt.Println("interface method call", time.Since(n))
}
