
package main

import "fmt"
/*类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：

type_name(expression)
type_name 为类型，expression 为表达式。

 */
/*Go 语言提供了另外一种数据类型即接口，
它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

 */

/* 定义接口 */
/*
type interface_name interface {
	method_name1 [return_type]
method_name2 [return_type]
method_name3 [return_type]
...
method_namen [return_type]
}

 */

/* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
//	/* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
//	/* 方法实现*/
//}
type Phone interface {

	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
//例子2
type Man interface {
	name() string  //有返回值的接口方法
	age() int
}

type Woman struct {
}

func (woman Woman) name() string {
	return "Jin Yawei"
}
func (woman Woman) age() int {
	return 23
}

type Men struct {
}

func ( men Men) name() string {
	return "liweibin"
}
func ( men Men) age() int {
	return 27
}
func main() {
	var sum int = 17
	var count int = 5
	var mean float32
//转类型
	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)


	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	//例子2
	var man Man;

	man = new(Woman)
	fmt.Println( man.name())
	fmt.Println( man.age())
	man = new(Men)
	fmt.Println( man.name())
	fmt.Println( man.age())
}