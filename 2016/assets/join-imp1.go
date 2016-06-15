package main

import (
	"fmt"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(join(s, ", "))
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
