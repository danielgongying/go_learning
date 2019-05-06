package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/afero"
	"io/ioutil"
	"time"
)

func main() {

	c := cache.New(3*time.Second,1*time.Minute)
	c.Set("name","daniel",cache.DefaultExpiration)
	name,isFound:=c.Get("name")
	if isFound {
		fmt.Println("name=",name)
	}else{
		fmt.Println("not found")

	}

	c.Delete("name")
	c.Add("name","mandy",cache.DefaultExpiration)
	c.DeleteExpired()
	newName,isFound:=c.Get("name")
	if isFound {
		fmt.Println("newname=",newName)
	}else{
		fmt.Println("not found")

	}


	fs:=afero.NewOsFs()
	file,err:=fs.Create("/Users/daniel/Desktop/test/test.txt")
	if err!=nil	 {
		fmt.Println(err)

	}
	fmt.Println(file.Name())
	byte,_:=file.WriteString("hello world")
	fmt.Println(byte)
	b,_:=ioutil.ReadFile(file.Name())
	fmt.Println(string(b))
    fs.Remove(file.Name())
	fs.RemoveAll("/Users/daniel/Desktop/test/")
	//fs.Remove("/Users/daniel/desktop/test/test.rtf")


}

