//dup的前两个版本以"流”模式读取输入，并根据需要拆分成多个行。
// 理论上，这些程序可以处理任意数量的输入数据。
// 还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			//进入下一次迭代
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
