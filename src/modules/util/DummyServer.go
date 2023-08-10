package util

import (
	log "bits/modules/common/log"
	"net/http"
)


type DummyServer struct {
	logger *log.CommonLogger
}

func (d DummyServer) Execute(port string) (bool, string) {

	d.logger = log.GetLogger()

	d.logger.Trace("Listen Dummy Server : " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil{
		return false, err.Error()
	}

	return true, ""
}
