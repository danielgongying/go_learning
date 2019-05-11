package main

import (
	"fmt"
	"strings"
)

func hello(s string) string {
	return ("hello" + s)

}
func main() {
	fmt.Println(hello("world"))
	a := strings.Contains("hello", "el")
	fmt.Println(a)
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true
	/*是否含有任意字符*/
	fmt.Println(strings.ContainsAny("team", "i"))       //false
	fmt.Println(strings.ContainsAny("failure", "wwwi")) //true
	fmt.Println(strings.ContainsAny("foo", ""))         //false
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) //["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         //[" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            //[""]
	fmt.Println(strings.HasSuffix("astaxie", "as"))                        //false
	fmt.Println(strings.HasSuffix("astaxie", "xie"))                       //true

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
	fmt.Println(strings.Title("her royal highness"))
	//Her Royal Highness
	fmt.Println(strings.ToUpper("Gopher"))
	//GOPHER

	read := strings.NewReader("I am asta谢")
	var b []byte
	fmt.Println(read.Len()) //12
	b = make([]byte, 8)
	n, err := read.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b) //[73 32 97 109 32 97 115 116]
	fmt.Println(n) //8
	n, err = read.ReadAt(b, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b) //[109 32 97 115 116 97 232 176]
	fmt.Println(n) //8
	bt, err := read.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bt) //97
	rn, size, err := read.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rn)   //35874
	fmt.Println(size) //3
}
