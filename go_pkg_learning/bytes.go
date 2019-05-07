package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.Map(func(r rune) rune {
		return r + 1
	}, []byte("345671abc"))
	fmt.Println(string(b))
	fmt.Println(bytes.HasPrefix(b, []byte("456")))
	fmt.Println(bytes.HasSuffix(b, []byte("345")))
	fmt.Println(string(bytes.Repeat(b, 3)))
	fmt.Println(string(bytes.Title(b)))
	fmt.Println(string(bytes.ToUpper(b)))
	for i, c := range bytes.Split(b, []byte{'b'}) {
		fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
	}
	fmt.Println(string(bytes.Replace(b, []byte("456"), []byte("abc"), 1)))

	var b3 *bytes.Buffer
	fmt.Println(b3.String())
	b3 = bytes.NewBuffer([]byte("123"))
	fmt.Println(b3.String())

	b4 := bytes.NewBuffer([]byte("123,456"))
	line, err := b4.ReadBytes(',')
	fmt.Println(string(line))
	fmt.Println(err)

	b5 := bytes.NewBuffer([]byte("123*456"))
	line1, err := b5.ReadString('*')
	fmt.Println(line1, err)
	line2, err := b5.ReadBytes('5')
	fmt.Println(string(line2))

	b6 := bytes.NewBuffer(nil)
	b6.WriteByte('1')
	b6.WriteByte('2')
	b6.WriteByte('3')
	fmt.Println(string(b6.Bytes()))

	b7 := bytes.NewBuffer(nil)
	b7.WriteString("你好世界")
	fmt.Println(string(b7.Bytes()))

}
