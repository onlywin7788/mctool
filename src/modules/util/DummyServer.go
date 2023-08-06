package util

import (
	log "bits/modules/common/log"
	"net/http"
)


type DummyServer struct {
}

func (dummyServer DummyServer) Execute(port string) (bool, string) {

	logger := log.CommonLogger{}
	logger.Trace("Listen Dummy Server : " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil{
		return false, err.Error()
	}

	return true, ""
}
