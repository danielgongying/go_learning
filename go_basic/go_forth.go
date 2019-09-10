package main

import (
	"fmt"
)
//切片有点像可变数组，可以追加元素，在追加时可能使切片的容量增大
/*你可以声明一个未指定大小的数组来定义切片：

var identifier []type
	切片不需要说明长度。

或使用make()函数来创建切片:

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)

 */
func main() {

	/*var numbers = make([]int,3,5)

	printSlice(numbers)

	 */

	/* 创建切片,数组需要告知数组元素个数 */
	numbers := []int{0,1,2,3,4,5,6,7,8}

	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

//func printSlice(x []int){
//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}