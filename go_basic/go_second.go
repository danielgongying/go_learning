package go_basic


import "fmt"
/*//声明数组
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//该实例与上面的实例是一样的，虽然没有设置数组的大小。

balance[4] = 50.0

 */
func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}


	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	 n := max(3,5)
	 fmt.Println(n)
	 x,y:=swap("nihao","world")
	 fmt.Println(x,y)
}
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}
//Go 函数可以返回多个值，例如：

func swap(x, y string) (string, string) {
	return y, x
}