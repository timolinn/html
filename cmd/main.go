package main

import (
	"bytes"
	"fmt"
	"html-parser"
)

func main() {
	r := bytes.NewReader([]byte("<h1>Hello Parser</h1>"))
	parsed := html.Parse(r)
	fmt.Println(string(parsed))
}
