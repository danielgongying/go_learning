package main

import "fmt"
//闭包 返回值是一个函数
func getLimit() func() int {
	limit := 10
	return func() int {
		limit -= 1
		return limit
	}
}

func main() {
	limit := getLimit()
	fmt.Println(limit()) // 9
	fmt.Println(limit()) // 8

	limit2 := getLimit()
	fmt.Println(limit2()) // 9
	fmt.Println(limit2()) // 8

	fmt.Println(limit()) // 7

}