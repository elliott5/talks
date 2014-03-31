// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func waiter(i int, done chan struct{}) {
	fmt.Println(i, "waiting...")
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "done!")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	for i := 0; i < 4; i++ {
		go waiter(i, done)
	}
	for i := 0; i < 4; i++ {
		<-done
	}
}

// endmain OMIT
