package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	src := strings.NewReader("int num = 1;")
	var s scanner.Scanner

	s.Init(src)

	s.Scan()

	//this will print the next token "int "to stdout
	fmt.Println(s.TokenText())
	fmt.Println(s.Pos().Column)
}
