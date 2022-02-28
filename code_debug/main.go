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
}
