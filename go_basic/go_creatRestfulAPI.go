package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

type Article struct {
	Title string`json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`

}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles:= Articles{
		Article{Title:"test title",Desc:"test desc",Content:"this is a test"},
		Article{Title:"test title",Desc:"test desc",Content:"this is a test"},
		Article{Title:"test title",Desc:"test desc",Content:"this is a test"},
		Article{Title:"test title",Desc:"test desc",Content:"this is a test"},
		Article{Title:"test title",Desc:"test desc",Content:"this is a test"},
	}
	//转成json格式
	json.NewEncoder(w).Encode(articles)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"HomePage endpoint hit")

}
func handleRequest() {

	//http://localhost:8081
	http.HandleFunc("/",homePage)
	//http://localhost:8081/articles
	http.HandleFunc("/articles",allArticles)

	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main() {
	handleRequest()
}