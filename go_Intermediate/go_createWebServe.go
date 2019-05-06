package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func creatWeb(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"hello,web")

}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})
	http.HandleFunc("/test", creatWeb)

	log.Fatal(http.ListenAndServe(":8081", nil))

}