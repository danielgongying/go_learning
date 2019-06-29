package main

import (
	"github.com/go-redis/redis"
	"html/template"
	"net/http"

)
var client *redis.Client
var templates *template.Template

func index(w http.ResponseWriter,r *http.Request)  {
	comments,err:=client.LRange("comments",0,10).Result()

	if err != nil {
		return
	}
	templates.ExecuteTemplate(w,"index.html",comments)
}
func main() {
	client = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/index.html"))
	http.HandleFunc("/",index)
	http.ListenAndServe(":8080",nil)
}


