package util

import (
	"net"
)

type FirewallChecker struct {
}


func (firewallChecker FirewallChecker) Execute(addr string) (bool, string) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false, err.Error()
	}
	defer conn.Close()

	return true, ""
}
