package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("www.baidu.com"))
	fmt.Println(net.ParseIP("127.168.1.2"))
	//	<nil>
	//	127.168.1.2
}
