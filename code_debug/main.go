package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	fmt.Println(net.ParseIP("http://127.0.0.1:8080"))
	fmt.Println(net.ParseIP("127.0.0.1:8080"))
	fmt.Println(net.ParseIP("www.baidu.com"))
	fmt.Println(strconv.Atoi("12h"))
	fmt.Println(strconv.FormatUint(uint64(123456), 16))
	var ch chan int
	ch = nil
	select {
	case a := <-ch:
		fmt.Println(a)
	default:

	}
	fmt.Println("over")
}
