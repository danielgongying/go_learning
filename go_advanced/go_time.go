package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Write([]byte("hello world"))

	t1 := time.Since(t)
	fmt.Print(t1)


}
func main() {
	http.HandleFunc("/",hello)
	http.ListenAndServe(":8080",nil)
	//r := chi.NewRouter()
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	//http.ListenAndServe(":8081",r)



}
