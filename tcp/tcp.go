package tcp

import (
	"fmt"
	"net"
)

const (
	lib = iota
	myImpl
)

var mode = lib

func NewTCPConnection(host string, port int) (net.Conn, error) {
	if mode == lib || true {
		return net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	}
	return net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
}
