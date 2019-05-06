package main

import (
	"fmt"
	"sync"
)
//
//func myFunc() {
//	fmt.Println("Inside my goroutine")
//}
//
//func main() {
//	fmt.Println("Hello World")
//	//main函数实际上在goroutine 获得执行机会之前终止。
//	go myFunc()
//	fmt.Println("Finished Execution")
//}
func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}

func main() {
	fmt.Println("Hello World")
//运用WaitGroup阻塞main线程，直到其中的任何goroutine WaitGroup成功执行。
//和运用channel达到一样的效果
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}