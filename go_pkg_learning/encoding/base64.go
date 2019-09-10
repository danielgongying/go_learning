package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"

	//"bytes"
	"encoding/base64"
	"fmt"
)

func encode()  {
	// 要编码的字符串
	src := "this is a test string."
	//编码成base64string
	dst:=base64.StdEncoding.EncodeToString([]byte(src))
	fmt.Println(string(dst))

	//// 计算编码缓冲区需要的长度
	//l := base64.StdEncoding.EncodedLen(len(src))
	//// 创建足够长度的缓冲区
	//dst := make([]byte, l)
	//base64.StdEncoding.Encode(dst,[]byte(src))
	//fmt.Println(string(dst))
	//
	//
	//解析string
	//dst,_=base64.StdEncoding.DecodeString(string(dst))
	//src = string(dst)
	//fmt.Println(string(dst))
	//print(src)

	data := []byte("123456")
	fmt.Printf("%x", md5.Sum(data))
	fmt.Printf("%x,", sha256.Sum256(data))
	fmt.Printf("%x,", sha512.Sum512(data))


}
func main() {
	encode()
	//bytes的运用

}