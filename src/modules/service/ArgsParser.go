/*****************************************

도움말 : -help
환경 체크 : -check -system
방화벽 점검 : -check_firewall=1 10.10.10.10:34952
임시 서버 기동(방화벽 체크용) : -server 8080
설정 템플릿 출력 : -template


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
	HELP int
	CHECK int
	SYSTEM_CHECK int
	FIREWALL_CHECK int
	FIREWALL_ADDR string
	TEMPLATE int
	DUMMYSERVER int
	DUMMYSERVER_PORT string
	DEBUG_SERVER int
	DEBUG_SERVER_PORT string
	DEBUG_CLIENT int
	DEBUG_CLIENT_ADDR string
}

func (argsParser ArgsParser) PrintHELP() string {

	helpContents :=
`
Usage:
-check -system                  Checking MicroStrategy Installation Enviroment (Linux Only)
-check -firewall IP:PORT        Firewall check   ex) -check firewall 10.10.10.10:34952
-server 34952                   Dummy Server Listen (Purpose checking firewall)
-template                       Printing Template Contents for MicroStrategy configiration


[Hidden usage / Deveopment-Only]
-debug_server PORT              Start remote Debugging Server (Windows - Linux Cross Debugging)
-debug_client IP:PORT           Client Debugging   ex) -check -system -debug_client 10.10.10.10:18080
`
	return helpContents
}

func (argsParser *ArgsParser) Parser(args []string) bool {

	// init
	argsParser.HELP = 0
	argsParser.TEMPLATE = 0
	argsParser.DUMMYSERVER = 0
	argsParser.DUMMYSERVER_PORT = ""
	argsParser.CHECK = 0
	argsParser.SYSTEM_CHECK = 0
	argsParser.FIREWALL_CHECK = 0
	argsParser.FIREWALL_ADDR = ""
	argsParser.DEBUG_SERVER = 0
	argsParser.DEBUG_SERVER_PORT = ""
	argsParser.DEBUG_CLIENT = 0
	argsParser.DEBUG_CLIENT_ADDR = ""

	logger := log.CommonLogger{}
	logger.Dummy()

	if args[0] == "-help" || args[0] == "-h" {
		argsParser.HELP = 1
	}

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
	if args[0] == "-server" {
		argsParser.DUMMYSERVER = 1
		argsParser.DUMMYSERVER_PORT = args[1]

	}
	if args[0] == "-debug_server"{
		argsParser.DEBUG_SERVER = 1
		argsParser.DEBUG_SERVER_PORT = args[1]
	}

	idx := 0
	for _, arg := range args {
        if arg == "-debug_client"{
			argsParser.DEBUG_CLIENT = 1
			argsParser.DEBUG_CLIENT_ADDR = args[idx+1]
		}
		idx++
    }

	return true
}
