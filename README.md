## 章节后的练习题汇总

- [练习题汇总](#章节后的练习题汇总)
    - [chapter1](#chapter1)
        - [1.1](#11)
        - [1.2](#12)
        - [1.3](#13)
        - [1.4](#14)
        - [1.5](#15)
        - [1.6](#16)
        - [1.7](#17)
        - [1.8](#18)
## chapter1
## 1.1
修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
~~~go
//每次循环迭代字符串s的内容都会更新。
// +=连接原字符串、空格和下个参数，产生新字符串，并把它赋值给s。
// s原来的内容已经不再使用，将在适当时机对它进行垃圾回收。
//如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用strings包的Join函数：
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
~~~
## 1.2
修改echo程序，使其打印每个参数的索引和值，每个一行。

## 1.3
做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
## 1.4
修改dup2，出现重复的行时打印文件名称。
## 1.5
修改前面的Lissajous程序里的调色板，由黑色改为绿色。我们可以用color.RGBA{0xRR, 0xGG, 0xBB, 0xff}来得到#RRGGBB这个色值，三个十六进制的字符串分别代表红、绿、蓝像素。
## 1.6
修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第三个参数，看看显示结果吧。
##1.7
函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
~~~go
/*
	练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，
	使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，
	避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "")
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "featch: %v\n", err)
			//终止进程
			os.Exit(1)
		}
		fmt.Printf("http status:%d\n",resp.StatusCode)

		out, err := os.Create("baidu-index.html")
		if err != nil {
			fmt.Fprintf(os.Stderr, "create file: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(out, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(b)
	}
}

~~~
