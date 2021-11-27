package ip

import (
	"fmt"
	"net"
)

const (
	lib = iota
	myImpl
)

var mode = lib
var modeV6 = lib

func NewIPConnection(network string, protcol int, addr string) (net.Conn, error) {
	if network == "ip" {
		if mode == lib {
			return net.Dial(fmt.Sprintf("%v:%v", network, protcol), addr)
		}
		return net.Dial(fmt.Sprintf("%v:%v", network, protcol), addr)
	}
	if network == "ip6" {
		if modeV6 == lib {
			return net.Dial(fmt.Sprintf("%v:%v", network, protcol), addr)
		}
		return net.Dial(fmt.Sprintf("%v:%v", network, protcol), addr)
	}
	return nil, fmt.Errorf("unsupported network: %v", network)
}
