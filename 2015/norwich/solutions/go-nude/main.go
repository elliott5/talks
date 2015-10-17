package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

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

	start := time.Now()
	singleThread(pth, fi)
	st := time.Since(start)
	fmt.Println("Single-thread time:", st, "\n")

	start = time.Now()
	goroutinePerFile(pth, fi)
	pf := time.Since(start)
	pfpct := float64(pf.Nanoseconds()) * 100 / float64(st.Nanoseconds())
	fmt.Println("Goroutine per file time:", pf,
		"Proportion of single-thread:", pfpct, "%\n")

	for nc := 1; nc <= 32; nc++ {
		start = time.Now()
		goroutinesPerCPU(pth, fi, nc)
		pc := time.Since(start)
		pcpct := float64(pc.Nanoseconds()) * 100 / float64(st.Nanoseconds())
		fmt.Println(nc, "goroutine(s) per CPU time:", pc,
			"Proportion of single-thread:", pcpct, "%\n")
		if pcpct <= pfpct {
			fmt.Println(nc,
				"goroutine(s) per CPU gives as good performance as per File")
			break
		}
	}
}

func singleThread(pth string, fi []os.FileInfo) {
	seen := 0
	arenude := 0
	for _, f := range fi {
		nam := f.Name()
		if strings.HasSuffix(strings.ToLower(nam), ".jpg") {
			seen++
			isNude, err := nude.IsNude(pth + nam)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(nam, "isNude=", isNude)
			if isNude {
				arenude++
			}
		}
	}
	fmt.Println("Proportion nude=", arenude, "/", seen, "=", float64(arenude)/float64(seen))
}

func goroutinePerFile(pth string, fi []os.FileInfo) {
	scores := make(chan bool)
	seen := 0
	for _, f := range fi {
		nam := f.Name()
		if strings.HasSuffix(strings.ToLower(nam), ".jpg") {
			seen++
			go func(nam string) { // a goroutine to process each file
				isNude, err := nude.IsNude(pth + nam)
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Println(nam, "isNude=", isNude)
				scores <- isNude
			}(nam)
		}
	}
	arenude := 0
	for i := 0; i < seen; i++ {
		wasNude := <-scores
		if wasNude {
			arenude++
		}
	}
	fmt.Println("Proportion nude=", arenude, "/", seen, "=", float64(arenude)/float64(seen))
}

func goroutinesPerCPU(pth string, fi []os.FileInfo, numGR int) {
	scores := make(chan bool)
	done := make(chan struct{})
	cpus := runtime.NumCPU() * numGR
	queue := make(chan string)
	go func() { // a goroutine to put the work to be done into the queue
		for _, f := range fi {
			nam := f.Name()
			if strings.HasSuffix(strings.ToLower(nam), ".jpg") {
				queue <- nam
			}
		}
		close(queue)
	}()
	for c := 0; c < cpus; c++ {
		go func() { // begin one worker goroutine for each cpu
			for nam := range queue {
				isNude, err := nude.IsNude(pth + nam)
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Println(nam, "isNude=", isNude)
				scores <- isNude
			}
			done <- struct{}{}
		}()
	}
	go func() { // a goroutine to wait for all the workers to be done
		// NOTE: sync.Waitgroup
		for c := 0; c < cpus; c++ {
			<-done
		}
		close(scores)
	}()
	seen := 0
	arenude := 0
	for wasNude := range scores {
		seen++
		if wasNude {
			arenude++
		}
	}
	fmt.Println("Proportion nude=", arenude, "/", seen, "=", float64(arenude)/float64(seen))
}
