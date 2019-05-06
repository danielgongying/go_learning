package main
import (
	"bufio"
	"fmt"
	"os"
)
//如果你只想从命令行中读取一个unicode字符，那么我建议你像这样使用bufio.ReadRune：
func main() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}
	scanner()
}
//无限地询问扫描输入并回显输入的内容。
func scanner() {
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
fmt.Println(scanner.Text())
}
}