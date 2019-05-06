package main
import (
	"encoding/json"
	"fmt"
)
func main ( ) {
	var jsonBlob = [ ] byte ( ` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] ` )
	type Animal struct {
		Name  string
		Order string
	}
	var animals [ ] Animal
	//json解析
	err := json. Unmarshal ( jsonBlob , & animals )
	if err != nil {
		fmt. Println ( "error:" , err )
	}

	fmt. Printf ( "%+v" , animals )
	a:=animals[0]
	fmt.Println(a)

}