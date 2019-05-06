package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(100)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {


	fmt.Println("Go Channel Tutorial")
	//如果通道未缓冲，则发送方将阻塞，直到接收方收到该值。
	//valueChannel := make(chan int)
	// 如果通道有缓冲区，则发送方仅阻塞，直到将值复制到缓冲区为止;
	// 如果缓冲区已满，则表示等待某个接收方检索到某个值。
	valueChannel := make(chan int,2)
	//go存在垃圾回收机制，所以无需显示关闭channel
	//defer close(valueChannel)

	go CalculateValue(valueChannel)
	//v := <-valueChannel
	go CalculateValue(valueChannel)
	values := <-valueChannel
	//fmt.Println(v,values)
	fmt.Println(values)
}