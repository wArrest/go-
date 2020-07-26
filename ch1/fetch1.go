package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "featch: %v\n", err)
			//终止进程
			os.Exit(1)
		}

		//函数从response中读取到全部内容,存在变量b中
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "featch: reading %s: %v\n", url, err)
			//终止进程
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
