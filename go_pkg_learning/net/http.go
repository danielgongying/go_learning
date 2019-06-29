package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServe(w http.ResponseWriter,req *http.Request)  {
	//io.WriteString(w,"hello world!\n")
	fmt.Fprint(w,"hello,.fsdfdsfds")
	req.Body=http.MaxBytesReader(w,req.Body,100)
	n,err:=io.Copy(ioutil.Discard,req.Body)
	if err == nil{
		fmt.Println("copy err")


	}
	if n != 20 {
		fmt.Println("io.Copy = ", n, ", want 20")
	}
	io.WriteString(w, "hello, world!hello, world!hello, world!hello, world!hello, world!\n")

}
func IndexPhpServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w,"<html><body>hello,this is a go page for " +
		"index.php</body></html>\n")

}
func NOFound(w http.ResponseWriter, req *http.Request)  {
	http.NotFound(w,req)

}
func error(w http.ResponseWriter,req *http.Request)  {
	http.Error(w,"this is a error",404)

}
func Redirect(w http.ResponseWriter,req *http.Request)  {
	http.Redirect(w,req,"url string",http.StatusFound)

}
func getHttp()  {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
func parseHttpVersion() {
	major, minor, ok := http.ParseHTTPVersion("HTTP/1.0")
	fmt.Println(major, minor, ok)
}

type Name struct {
	fname string
}
func GODemo(w http.ResponseWriter,req *http.Request)  {
	//name:= Name{"daniel"}
	//tpl,_ := template.ParseFiles("testgo.html")
	//tpl.Execute(w,name)
	t1, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")


}
func main() {
	getHttp()
	parseHttpVersion()
	http.HandleFunc("/hello",HelloServe)
	http.HandleFunc("/index.php",IndexPhpServer)
	http.HandleFunc("/",NOFound)
	http.HandleFunc("/error",error)
	http.HandleFunc("/url",Redirect)
	http.HandleFunc("/test",GODemo)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln(err)
	}





}

