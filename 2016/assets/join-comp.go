// +build OMIT

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const itters = 256
	base := make([]string, 0, itters*4)
	for i := 0; i < itters; i++ {
		base = append(base, []string{"foo", "bar", "baz", "Go London user group"}...)
	}

	s, l, n := base, itters, time.Now()
	for ; l > 0; l-- {
		_ = len(join(s, ", "))
	}
	fmt.Println("   my Join()", time.Since(n))

	s, l, n = base, itters, time.Now()
	for ; l > 0; l-- {
		_ = len(strings.Join(s, ", "))
	}
	fmt.Println("strings.Join()", time.Since(n))
}

func join(a []string, sep string) string {
	var ret string
	for k, v := range a {
		if k == 0 {
			ret = v
		} else {
			ret += sep + v
		}
	}
	return ret
}
