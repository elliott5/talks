package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// consume the service of github.com/go-kit/kit/examples/stringsvc1

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func main() {
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
}
