//但是，程序如何获取要处理的输入数据呢？
// 一些程序生成自己的数据，但通常情况下，输入来自于程序外部：文件、网络连接、其它程序的输出、敲键盘的用户、命令行参数或其它类似输入源。
// 首先是命令行参数。

//os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
//os.Args[0]，是命令本身的名字;其它的元素则是程序启动时传给它的参数

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("os.Args[0]:", os.Args[0])
	fmt.Println(s)
}
