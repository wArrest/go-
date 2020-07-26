// how to use for and slice
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//切片用法，类似python可以声明从哪个元素开始
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
