package main

import (
	"fmt"
	"unsafe"
)

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

//常量还可以用作枚举：

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
const (
	a2 = 1
	b2
	c2
	d2
)

func main() {
	fmt.Println("Hello, World!")
	var name string = "daniel"
	fmt.Println(name)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	v_name := "haha"
	fmt.Println(v_name)


	a = 9
	fmt.Println(a)
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

	n, i, h := numbers()
	print(n, i, h)

	//常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	//const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH

	fmt.Printf("面积为 : %d", area)
	fmt.Println(a1, b1, c1)
	println()
	//println(a, b, c)
	// b2、c2、d2没有初始化，使用上一行(即a)的值
	fmt.Println(b2) // 输出1
	fmt.Println(c2) // 输出1
	fmt.Println(d2) // 输出1

	//以下实例演示了各个算术运算符的用法：
	/*var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )
	*/

}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
