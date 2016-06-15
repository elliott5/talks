package main

import "fmt"

func main() {
	str := "Hello! 你好！"
	for key, value := range str {
		fmt.Println(key, value, string(value), string(str[key]))
        // NOTE: str[key]='x' is illegal
	}
}
