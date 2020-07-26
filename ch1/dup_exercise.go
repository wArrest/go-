//修改dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]map[string]int)
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

	for filename, lineCounts := range counts {
		for line, n := range lineCounts {
			if n > 1 {
				fmt.Printf("%s\n", filename)
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

}

//map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本）
func countLines(f *os.File, counts map[string]map[string]int) {

	if _, ok := counts[f.Name()]; !ok {
		counts[f.Name()] = make(map[string]int)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
