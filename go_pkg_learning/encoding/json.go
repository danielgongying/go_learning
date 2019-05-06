package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apsdehal/go-logger"
	"io"
	"log"
	"os"
	"strings"
)

/*json数据转成bytes.Buffer*/
func compact()  {
	dst := new(bytes.Buffer)
	src := []byte(`{
	 "name":"runoob",
    "alexa":10000,
    "sites": {
        "site1":"www.runoob.com",
        "site2":"m.runoob.com",
        "site3":"c.runoob.com"
    },
	"sites":[ "Google", "Runoob", "Taobao" ]
		}`)
	//json数据转成bytes.Buffer
	err := json.Compact(dst, src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dst)
	//转成string
	fmt.Println(dst.String())
}
/*json解析到struct*/
func decode()  {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`
	type Message struct {
		Name, Text  string
	}
	/*NewDecoder创建一个decoder实例*/
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text,)
	}
}
/*struct的数据转成json*/
func encode()  {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(group); err != nil {
		fmt.Printf("failed encoding to writer: %s", err)
	}
}
/*将对象v序列化为json格式[]byte*/
func marshal()  {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	//b, err := json.Marshal(group)
	/*对象v序列化为缩进格式json*/
	b, err := json.MarshalIndent(group, "##", "**")
	//实际应用中更多的可能为 空格，这里为了便于观察

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
	os.Stdout.Write(b)

}
/*反序列化JSON编码的data，并将结果存储到指向v的指针(对象struct)*/
func unmarshal()  {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

}
/*与Compact不同的是会将<， >， 和 & 编码成 \u003c, \u003e, \u0026，
以便于可以安全的嵌入到html的<script>标签中*/
func htmlEscape()  {
	dst := new(bytes.Buffer)
	src := []byte(`{
		"Name":"tony.shao",
		"Age":25,
		"Job":"Programmer<Escaping>"
		}`)
	json.HTMLEscape(dst, src)
	fmt.Println(dst)

}
/*缩进形式提那家到缩进的json中*/
func indent()  {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	dst := new(bytes.Buffer)
	src := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	json.Indent(dst, src, "#fsdf#", "") //实际应用中，可能更多的是使用空格来作为prefix和indent，这里主要为了方便观察
	json.Indent(dst, src, "##", "**") //实际应用中，可能更多的是使用空格来作为prefix和indent，这里主要为了方便观察
	fmt.Println(dst)
}
func main() {
	compact()
	decode()
	encode()
	logger, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	logger.Info("info")
	htmlEscape()
	indent()
	marshal()
	unmarshal()


}