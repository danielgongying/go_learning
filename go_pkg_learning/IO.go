package main
import (
	"fmt"
	"io"
	"os"
	"time"
)
func main() {
	file,_ := os.Open("/Users/daniel" +
		"/Desktop/昆咬视屏.png")
	toFile,_:=os.Create("/Users/" +
		"daniel/Desktop/test/image.png")
	w,err:=io.Copy(toFile,file)
	if err != nil {
		fmt.Println("success:",w)
	}
	time.Sleep(5*time.Second)
	os.Remove(toFile.Name())


}

