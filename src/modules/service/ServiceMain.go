package service

import (
	"fmt"
	"os"
	log "bits/modules/common/log"
	util "bits/modules/util"
	command "bits/modules/command"
)


type ServiceMain struct {
}

func (serviceMain ServiceMain) Execute() bool{

	logger := log.CommonLogger{}
	logger.Dummy()

	commander := command.Commander{}
	argsParser := ArgsParser{}

	args := os.Args[1:]
	argsParser.Parser(args)

	if argsParser.HELP == 1{

		printHELP()

	}else if argsParser.DEBUG_SERVER == 1 {

		commander.DebugServerListen(argsParser.DEBUG_SERVER_PORT)

	}else if argsParser.CHECK == 1 {

		if argsParser.SYSTEM_CHECK == 1 {
			CheckSystem(argsParser.DEBUG_CLIENT, argsParser.DEBUG_CLIENT_ADDR)
		}
		if argsParser.FIREWALL_CHECK == 1 {
			CheckFirewall(argsParser.FIREWALL_ADDR)
		}

	}else if argsParser.TEMPLATE == 1 {
	
		PrintTemplate()
	
	}else if argsParser.DUMMYSERVER == 1 {

		ListenDummyServer(argsParser.DUMMYSERVER_PORT)

	}else{

		logger.BasicPrint("\nInvalid Execute paramter. Please check with '-help'\n")

	}

	return true
}


func CheckSystem(debug_client int, debug_addr string) bool{
	
	logger := log.CommonLogger{}
	logger.Dummy()

	systemChecker := util.SystemChecker{}

	logger.Info("system checking")

	flag, errMsg := systemChecker.Execute(debug_client, debug_addr)

	if flag == false{
		logger.Error(errMsg)
	}
	return flag
}

func CheckFirewall(addr string) bool{
	
	logger := log.CommonLogger{}
	firewallChecker := util.FirewallChecker{}

	logger.Info("firewall checking : " + addr)
		
	flag, errMsg := firewallChecker.Execute(addr)

	if flag == true{
		logger.Info("Firewall opened.")
	} else{
		logger.Error(errMsg)
	}
	return true
}

func PrintTemplate() bool{
	
	logger := log.CommonLogger{}
	template := util.Template{}

	logger.Info("Print Template Configuration")
	template.Execute()

	return true
}

func ListenDummyServer(port string) bool{

	logger := log.CommonLogger{}
	dummyServer := util.DummyServer{}

	logger.Info("Dummy Server Listen - Port : " + port)
	flag, errMsg := dummyServer.Execute(port)
	if flag == false{
		logger.Error(errMsg)
	} 
	return true
}

func printHELP() bool{
	
	argsParser := ArgsParser{}
	content := argsParser.PrintHELP()
	fmt.Println(content)

	return true
}