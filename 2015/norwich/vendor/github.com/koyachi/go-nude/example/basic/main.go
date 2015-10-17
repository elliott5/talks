package main

import (
	"github.com/koyachi/go-nude"
	"log"
//	"syscall"
)

func main() {
//	syscall.UnzipFS("images.zip")
	imagePaths := []string{"../images/damita.jpg",
	 "../images/damita2.jpg",
	"../images/test2.jpg",
	 "../images/test6.jpg"}

for _,imagePath := range imagePaths {

	isNude, err := nude.IsNude(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	println(imagePath, "isNude=", isNude)
}}
