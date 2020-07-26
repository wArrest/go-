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
