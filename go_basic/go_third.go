package go_basic

import (
	"fmt"
)
//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	var a int = 10
	var ip *int
	fmt.Printf("变量的地址：%x\n", &a)
	fmt.Println("变量的地址：", &a)
	ip = &a
	fmt.Println("ip 变量存储的指针地址:", ip)
	fmt.Println("ip 变量存储的指针地址的值:", *ip)
	fmt.Println("ip 变量存储的指针地址的地址:", &ip)
	var ptr *int
	if (ptr != nil) {
		if (ip != nil) {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip不是空指针")
		}else {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip是空指针")
		}
	} else {
		if(ip != nil){
			fmt.Println("ptr是空指针")
			fmt.Println("ip不是空指针")
		}else{
			fmt.Println("ptr是空指针")
			fmt.Println("ip是空指针")
		}
	}
	/*  自学的时候想到能不能使用 switch 优化 for 繁琐的写法，但是发现 case 匹配到后会自动跳出 switch。
	    查了一下 select 等方法发现并不适用， 最后发现了 fallthrough 可以很好的用在这里（不过要注意 fallthrough 存在的位置，避免产生逻辑混乱）  */
	switch {
	case ptr != nil:
		fmt.Println("ptr不是空指针")
		fallthrough
	case ptr == nil:
		fmt.Println("ptr是空指针")
		fallthrough
	case ip != nil:
		fmt.Println("ip不是空指针")
	default:
		fmt.Println("ip是空指针")
	}



	// 创建一个新的结构体

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})


	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

