package main

import "fmt"

func main() {
	chn := make(chan string, 2)
	go func() {
		chn <- "Alpha"
		chn <- "Beta"
		chn <- "Gamma"
		chn <- "Delta"
		close(chn) // what happens when you comment this out?
	}()
	value, ok := <-chn
	fmt.Println(value, ok)
	for greek := range chn {
		fmt.Println(greek)
	}
	value, ok = <-chn
	fmt.Println(value, ok)
}
