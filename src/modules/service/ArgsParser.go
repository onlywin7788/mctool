/*****************************************

환경 체크 : -check -type=[all/rpm/config/resource]
방화벽 점검 : -check_firewall=1 10.10.10.10:34952

[디버깅모드 추가]
디버깅 서버 : -debug_server
클라이언트 : {기존 파라미터} -debug_client

*****************************************/

package service

import (
	_ "fmt"
	log "bits/modules/common/log"
)

type ArgsParser struct {
	CHECK int
	SYSTEM_CHECK int
	FIREWALL_CHECK int
	FIREWALL_ADDR string
	TEMPLATE int
	DEBUG_SERVER int
	DEBUG_CLIENT int
}


func (argsParser *ArgsParser) Parser(args []string) bool {

	// init
	argsParser.TEMPLATE = 0
	argsParser.CHECK = 0
	argsParser.SYSTEM_CHECK = 0
	argsParser.FIREWALL_CHECK = 0
	argsParser.FIREWALL_ADDR = ""
	argsParser.DEBUG_SERVER = 0
	argsParser.DEBUG_CLIENT = 0

	logger := log.CommonLogger{}
	logger.Dummy()

//	fmt.Println(args)

	if args[0] == "-check" {
		
		argsParser.CHECK = 1

		if args[1] == "-system" {
			argsParser.SYSTEM_CHECK = 1
		}
		if args[1] == "-firewall" {
			argsParser.FIREWALL_CHECK = 1
			argsParser.FIREWALL_ADDR = args[2]
		}
	}
	if args[0] == "-template" {
		argsParser.TEMPLATE = 1
	}
	if args[0] == "-debug_server"{
		argsParser.DEBUG_SERVER = 1
	}

	for _, arg := range args {
        if arg == "-debug_client"{
			argsParser.DEBUG_CLIENT = 1
		}
    }

	return true
}
