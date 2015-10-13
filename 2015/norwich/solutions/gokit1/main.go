package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func ucase(s string) (string, error) {
	req := uppercaseRequest{S: s}
	j, e0 := json.Marshal(req)
	if e0 != nil {
		return "", e0
	}
	r, e1 := http.Post("http://localhost:8080/uppercase", "application/json",
		bytes.NewReader(j))
	if e1 != nil {
		return "", e1
	}
	b, e2 := ioutil.ReadAll(r.Body)
	if e2 != nil {
		return "", e2
	}
	res := &uppercaseResponse{}
	e3 := json.Unmarshal(b, &res)
	if e3 != nil {
		return "", e3
	}
	if len(res.Err) > 0 {
		e3 = errors.New(res.Err)
	}
	return res.V, e3
}

func count(s string) (int, error) {
	req := countRequest{S: s}
	j, e0 := json.Marshal(req)
	if e0 != nil {
		return 0, e0
	}
	r, e1 := http.Post("http://localhost:8080/count", "application/json",
		bytes.NewReader(j))
	if e1 != nil {
		return 0, e1
	}
	b, e2 := ioutil.ReadAll(r.Body)
	if e2 != nil {
		return 0, e2
	}
	res := &countResponse{}
	e3 := json.Unmarshal(b, &res)
	if e3 != nil {
		return 0, e3
	}
	return res.V, nil
}

func main() {
	u, e := ucase("zombie")
	fmt.Println("ucase", u, e)
	c, e2 := count("zombie")
	fmt.Println("count", c, e2)
}
