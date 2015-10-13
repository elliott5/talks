package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/koyachi/go-nude"
)

func main() {

	pth := "./rubens/Peter Paul Rubens - Wikipedia, the free encyclopedia_files/"

	dir, err := os.Open(pth)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	seen := 0
	arenude := 0
	scores := make(chan bool, len(fi))
	for _, f := range fi {
		nam := f.Name()
		if strings.HasSuffix(nam, ".jpg") {
			seen++
			go func(nam string) {
				isNude, err := nude.IsNude(pth + nam)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(nam, "isNude=", isNude)
				scores <- isNude
			}(nam)
		}
	}
	for i := 0; i < seen; i++ {
		wasNude := <-scores
		if wasNude {
			arenude++
		}
	}
	fmt.Println("Proportion nude=", arenude, "/", seen, "=", float64(arenude)/float64(seen))
}
