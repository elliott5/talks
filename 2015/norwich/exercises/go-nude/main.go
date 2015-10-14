package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/koyachi/go-nude"
)

func main() {
	sl := string(os.PathSeparator)
	pth := "." + sl + "rubens" + sl + "Peter Paul Rubens - Wikipedia, the free encyclopedia_files" + sl

	dir, err := os.Open(pth)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fi {
		nam := f.Name()
		if strings.HasSuffix(strings.ToLower(nam), ".jpg") {
			isNude, err := nude.IsNude(pth + nam)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(nam, "isNude=", isNude)
		}
	}
}
