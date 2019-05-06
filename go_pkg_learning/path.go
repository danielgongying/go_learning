package main

import (
	"fmt"
	"path"

)

func join()  {
	fmt.Println("join start")
	fmt.Println(path.Join("a", "b", "c")) // a/b/c
	fmt.Println(path.Join("a", "", "c"))  // a/c
	fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c

}
func clean()  {
	fmt.Println(path.Clean("a/c"))// a/c
	fmt.Println(path.Clean("a//c"))//以一个/代替// , a/c
	fmt.Println(path.Clean("a/c/."))//清除. , a/c
	fmt.Println(path.Clean("a/c/b/.."))// 清除内部..以及前面的元素b; a/c
	fmt.Println(path.Clean("111/../a/c"))// 清除内部..以及前面的元素111; a/c
	fmt.Println(path.Clean("/../a/b/../././/c"))// 清除/..开头,..以及前面的元素b; /a/c
}
func match()  {
	fmt.Println(path.Match("*","alll")) //true nil
	fmt.Println(path.Match("*","a/lll")) //false nil
	fmt.Println(path.Match("?","alll")) //false nil
	fmt.Println(path.Match("?","a")) //true nil
	fmt.Println(path.Match("a","a")) //true nil
	fmt.Println(path.Match("\\a","a")) //true nil

}
func dir()  {
	fmt.Println(path.Dir("/a/b/c")) // /a/b
	fmt.Println(path.Dir("")) // .

}
/*返回扩展名*/
func ext()  {
	fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
	fmt.Println(path.Ext("/a/b/c/bar"))	// ""
	fmt.Println(path.Ext("/a/b/c/bar.txt"))	// ""
}
/*是否绝对路径*/
func isAbs()  {
	fmt.Println(path.IsAbs("/home/zzz/go.pdf")) // true
	fmt.Println(path.IsAbs("home/zzz/go.pdf"))  // false

}
/*分离路径中的文件目录和文件*/
func split()  {
	fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css
	fmt.Println(path.Split("static/1.txt"))	// "" static

}
func main() {


	clean()
	join()
	match()
	dir()
	ext()
	isAbs()
	split()
}