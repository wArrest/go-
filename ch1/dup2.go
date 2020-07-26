//使用os.Open打开各个具名文件，并操作它们。

//// Dup2 prints the count and text of lines that appear more than once
//// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				//todo 
				fmt.Fprintf(os.Stderr, arg+": %v\n", err)
				//直接跳到下一个迭代
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

//map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本）
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
