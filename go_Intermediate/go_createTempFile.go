package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	fmt.Println(tmpfile.Name())

	//
	//dir, err := ioutil.TempDir("", "example")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer os.RemoveAll(dir) // clean up
	//
	//fmt.Println(dir)

}