package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
	Name string `json:"name"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
	Name string `json:"name"`

}

func main() {

	author := Author{Sales: 3, Age: 25, Developer: true ,Name:"daniel"}
	book := Book{Title: "Learning Concurrency in Python", Author: author,Name:"how to play go"}
	//对象转成json
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println(string(byteArray))
}