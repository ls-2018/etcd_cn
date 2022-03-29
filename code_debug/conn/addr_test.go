package conn

import (
	"fmt"
	"strings"
	"testing"
)

func TestConn(t *testing.T) {
	res := SubCommand([]string{"zsh", "-c", fmt.Sprintf("lsof -itcp -n|grep '%s->'", "53444")})
	pid := SubCommand([]string{"zsh", "-c", fmt.Sprintf("echo '%s'| awk '{print $2}'", res)})
	commandRes := SubCommand([]string{"zsh", "-c", fmt.Sprintf("ps -ef|grep -v grep |grep '%s'", pid)})
	command := SubCommand([]string{"zsh", "-c", fmt.Sprintf("echo '%s'| awk '{$1=$2=$3=$4=$5=$6=$7=\"\"; print $0}'", commandRes)})
	fmt.Println(strings.Trim(command, " "))
}
