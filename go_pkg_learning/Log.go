package main

import (
	"fmt"
	"github.com/apsdehal/go-logger"
	"log"
	"os"

)

func main() {
	//log.Fatal("print and os.exit")
	//log.Fatalf("print and os.exit again")
	//log.Fatalln("bye")
	fmt.Println(log.Flags())

	file, err := os.OpenFile("/Users/daniel/Desktop/demo.txt", os.O_WRONLY, 0666)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	l := log.New(file, "logger", log.Ldate)

	l.Println("log to file sample")	//logger2013/03/10 log to file sample

	logger, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	logger.Info("info")
	logger.Warning("warning")
	logger.Error("error")
	//logger.Fatal("bye")
	logger.Debug("debug")

}



