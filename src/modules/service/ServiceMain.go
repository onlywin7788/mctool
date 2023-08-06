package service

import (
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

	if argsParser.DEBUG_SERVER == 1 {
		commander.DebugServerListen()
	}

	if argsParser.CHECK == 1 {

		if argsParser.SYSTEM_CHECK == 1 {
			CheckSystem(argsParser.DEBUG_CLIENT)
		}
		if argsParser.FIREWALL_CHECK == 1 {
			CheckFirewall(argsParser.FIREWALL_ADDR)
		}
	}

	if argsParser.TEMPLATE == 1 {

	}

	return true
}


func CheckSystem(debug_client int) bool{
	
	logger := log.CommonLogger{}
	logger.Dummy()

	systemChecker := util.SystemChecker{}

	logger.Info("system checking")

	flag, errMsg := systemChecker.Check(debug_client)

	if flag == false{
		logger.Error(errMsg)
	}
	return flag
}

func CheckFirewall(addr string) bool{
	
	logger := log.CommonLogger{}
	firewallChecker := util.FirewallChecker{}

	logger.Info("firewall checking : " + addr)
		
	flag, errMsg := firewallChecker.Check(addr)

	if flag == true{
		logger.Info("Firewall opened.")
	} else{
		logger.Error(errMsg)
	}
	return true
}

func priontTemplate(addr string) bool{
	
	logger := log.CommonLogger{}
	firewallChecker := util.FirewallChecker{}

	logger.Info("Print Template Configuration")

	
	return true
}