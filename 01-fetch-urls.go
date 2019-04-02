package main

import (
	"fmt"
	"os"
)

func main() {


	urls := []string{
		"www.baidu.com",
		"www.golang.com",
		"www.facebooke.cn",
	}

	ch := make(chan string)
	for _, url := range urls {
		go fetchOne(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	os.Exit(0)
}

func fetchOne(url string, result chan<- string) {

	result <- fmt.Sprintf("Fetching %s", url)
}
