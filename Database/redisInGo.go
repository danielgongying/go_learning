package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gobuffalo/packr"
	"html/template"
	"io"
	"net/http"

)
var client *redis.Client
var templates *template.Template

func index(w http.ResponseWriter,r *http.Request)  {
	comments,err:=client.LRange("comments",0,10).Result()

	if err != nil {
		return
	}
	// set up a new box by giving it a (relative) path to a folder on disk:
	box := packr.NewBox("./templates")

	// Get the string representation of a file, or an error if it doesn't exist:
	html, err := box.FindString("index.html")
	templates.ExecuteTemplate(w,html,comments)
}

func IndexPhpServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w,"<html><body>hello,this is a go page for " +
		"index.php</body></html>\n")

}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
	})
	// set up a new box by giving it a (relative) path to a folder on disk:
	box := packr.NewBox("./templates")
	fmt.Println(box)
	// Get the string representation of a file, or an error if it doesn't exist:
	html, _ := box.FindString("index.html")

	print(html)
	templates = template.Must(template.ParseFiles(html))
	http.HandleFunc("/",index)
	http.HandleFunc("/php",IndexPhpServer)
	http.ListenAndServe(":8080",nil)

}


