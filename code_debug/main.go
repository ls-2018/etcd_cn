package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	math_bits "math/bits"
	"net"
	"strconv"
	"time"

	"github.com/ls-2018/etcd_cn/offical/etcdserverpb"
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

	a := `{"header":{"ID":7587861231285799685},"put":{"key":"YQ==","value":"Yg=="}}`
	b := `{"ID":7587861231285799684,"Method":"PUT","Path":"/0/version","Val":"3.5.0","Dir":false,"PrevValue":"","PrevIndex":0,"Expiration":0,"Wait":false,"Since":0,"Recursive":false,"Sorted":false,"Quorum":false,"Time":0,"Stream":false}`
	fmt.Println(json.Unmarshal([]byte(a), &etcdserverpb.InternalRaftRequest{}))
	fmt.Println(json.Unmarshal([]byte(b), &etcdserverpb.InternalRaftRequest{}))
	var c time.Time
	fmt.Println(c.IsZero())
}
