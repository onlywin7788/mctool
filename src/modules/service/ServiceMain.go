package service

import (
	"os"
	log "bits/modules/common/log"
	util "bits/modules/util"
	command "bits/modules/command"
)


type ServiceMain struct {
	logger *log.CommonLogger
}


func (s ServiceMain) Execute() bool{

	s.logger = log.GetLogger()

	commander := command.Commander{}
	argsParser := ArgsParser{}

	args := os.Args[1:]
	argsParser.Parser(args)

	s.logger.SetLogLevel(argsParser.LOGLEVEL)

	if argsParser.HELP == 1{

		s.printHELP()

	}else if argsParser.DEBUG_SERVER == 1 {

		commander.DebugServerListen(argsParser.DEBUG_SERVER_PORT)

	}else if argsParser.CHECK == 1 {

		if argsParser.SYSTEM_CHECK == 1 {
			s.checkSystem(argsParser.DEBUG_CLIENT, argsParser.DEBUG_CLIENT_ADDR)
		}
		if argsParser.FIREWALL_CHECK == 1 {
			s.checkFirewall(argsParser.FIREWALL_ADDR)
		}

	}else if argsParser.TEMPLATE == 1 {
	
		s.printTemplate()
	
	}else if argsParser.DUMMYSERVER == 1 {

		s.listenDummyServer(argsParser.DUMMYSERVER_PORT)

	}else{

		s.logger.BasicPrint("\nInvalid Execute paramter. Please check with '-help'\n")

	}

	return true
}


func (s ServiceMain)checkSystem(debug_client int, debug_addr string) bool{
	
	systemChecker := util.SystemChecker{}

	s.logger.Info("system checking")

	flag, errMsg := systemChecker.Execute(debug_client, debug_addr)

	if flag == false{
		s.logger.Error(errMsg)
	}
	return flag
}

func (s ServiceMain)checkFirewall(addr string) bool{
	
	firewallChecker := util.FirewallChecker{}

	s.logger.Info("firewall checking : " + addr)
		
	flag, errMsg := firewallChecker.Execute(addr)

	if flag == true{
		s.logger.Info("Firewall opened.")
	} else{
		s.logger.Error(errMsg)
	}
	return true
}

func (s ServiceMain)printTemplate() bool{
	
	template := util.Template{}

	s.logger.Info("Print Template Configuration")
	template.Execute()

	return true
}

func (s ServiceMain)listenDummyServer(port string) bool{

	dummyServer := util.DummyServer{}

	s.logger.Info("Dummy Server Listen - Port : " + port)
	flag, errMsg := dummyServer.Execute(port)
	if flag == false{
		s.logger.Error(errMsg)
	} 
	return true
}

func (s ServiceMain)printHELP() bool{
	
	argsParser := ArgsParser{}
	content := argsParser.PrintHELP()
	s.logger.BasicPrint(content)

	return true
}