package conn

import (
	"fmt"
	"net"
)

func PrintConn(line string, c net.Conn) {
	fmt.Println(fmt.Sprintf("%s----->localAddr:%s   RemoteAddr:%s", line, c.LocalAddr().String(), c.RemoteAddr().String()))
}
