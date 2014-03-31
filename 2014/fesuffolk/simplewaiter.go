// +build OMIT

package main

import (
	"fmt"
	"time"
)

func waiter(i int) {
	fmt.Println(i, "waiting...")
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(i, "done!")
}

func main() {
	for i := 4; i > 0; i-- {
		go waiter(i)
	}
	time.Sleep(5 * time.Second)
}
