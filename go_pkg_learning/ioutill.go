package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//b,err:=ioutil.ReadFile("/Users/daniel/Desktop/直播需求.pages")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(b))
	b,err:=ioutil.ReadDir("/Users/daniel/Desktop/test")
	if err != nil {
		fmt.Println(err)
	}
	for i,v := range b{
		fmt.Println(i,"=",v.Name(),v.Size())
		os.Remove(v.Name())
	}


	s := strings.NewReader("hello world")
	w,e:=ioutil.ReadAll(s)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(w))
	r:=ioutil.NopCloser(s)
	defer r.Close()

	f,er:=ioutil.TempDir("/Users/daniel/Desktop/test",
		"temp")
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(f)
	os.RemoveAll(f)

	f2,er2:=ioutil.TempFile("/Users/daniel/Desktop/test",
		"temp")
	if er2 != nil {
		fmt.Println(er2)

	}
	f2.WriteString("i am here")
	b2,e2:=ioutil.ReadFile(f2.Name())
	if e2 != nil {
		fmt.Println(e2)

	}
	fmt.Println(string(b2))

  os.Remove(f2.Name())


}

