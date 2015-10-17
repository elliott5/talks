package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go servermain() // create a very simple go-kit server
	time.Sleep(time.Second)

	// just a simple "curl" to get you started...
	r, e1 := http.Post("http://localhost:8080/count", "application/json",
		bytes.NewReader([]byte(`{"s":"hello, world"}`)))
	if e1 != nil {
		log.Fatal(e1)
	}
	b, e2 := ioutil.ReadAll(r.Body)
	if e2 != nil {
		log.Fatal(e2)
	}
	fmt.Println(string(b))

	for { // infinite loop so that curl commands will work
	}
}
