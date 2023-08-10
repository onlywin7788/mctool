package util

import (
	"net"
	log "bits/modules/common/log"
)

type FirewallChecker struct {
	logger *log.CommonLogger
}


func (f FirewallChecker) Execute(addr string) (bool, string) {

	f.logger = log.GetLogger()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false, err.Error()
	}
	defer conn.Close()

	return true, ""
}
