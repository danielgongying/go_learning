package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func testGet() {
	resp,err:=http.Get("http://www.720fy.com/")
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body))
	}else{
		fmt.Println(err.Error())
	}


}
func test1()  {
	handler := func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "<html><body>Hello World!</body></html>")
		w.Write([]byte("<html><body>Hello World!</body></html>"))
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
func test2(){
	res, err := http.Get("http://www.720fy.com/api/public/" +
		"index.php?service=Live.checkLive")
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

func main() {

	//testGet()
	//test1()
	test2()



	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}

