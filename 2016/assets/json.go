package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Response struct {
	Page   int
	Fruits []string
}

func main() {
	j := `{"Page": 42, "Fruits": ["banana", "pineapple"]}`
	fmt.Println(j)
	res := &Response{}
	json.Unmarshal([]byte(j), &res)
	fmt.Printf("%d %v\n", res.Page, res.Fruits)
	enc := xml.NewEncoder(os.Stdout)
	enc.Encode(res)
}
