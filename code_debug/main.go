package main

import (
	"crypto/md5"
	"fmt"
	math_bits "math/bits"
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
	var x uint64
	x = 1
	fmt.Println((math_bits.Len64(x|1) + 6) / 7)
	fmt.Println("over")
	hash := md5.New()
	hash.Write([]byte("hello"))
	fmt.Println(fmt.Sprintf("%x", hash.Sum(nil)))
}
