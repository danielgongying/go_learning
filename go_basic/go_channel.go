package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	//【0，10）的随机数
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	//value发到channel
	values <- value
}

func main() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go CalculateValue(values)
//从channel中取值
	value := <-values
	fmt.Println("value=",value)
}