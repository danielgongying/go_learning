package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
func main() {
	file,_ := os.Open("/Users/daniel" +
		"/Desktop/map.png")
	toFile,_:=os.Create("/Users/" +
		"daniel/Desktop/image.png")
	w,err:=io.Copy(toFile,file)
	if err != nil {
		fmt.Println("success:",w)
	}
	time.Sleep(5*time.Second)
	os.Remove(toFile.Name())




}

