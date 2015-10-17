package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func acase(which, s string) (string, error) {
	req := uppercaseRequest{S: s}
	j, e0 := json.Marshal(req)
	if e0 != nil {
		return "", e0
	}
	r, e1 := http.Post("http://localhost:8080/"+which+"case", "application/json",
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
	go servermain() // start the local server
	time.Sleep(time.Second)

	msg := "Hello Norwich!"

	c, e := count(msg)
	fmt.Println("count", c, e)
	u, e := acase("upper", msg)
	fmt.Println("u/case", u, e)
	l, e := acase("lower", msg)
	fmt.Println("l/case", l, e)
}
