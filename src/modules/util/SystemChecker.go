package util

import (
	"fmt"
	"strings"
	log "bits/modules/common/log"
	command "bits/modules/command"
	"encoding/json"
)

type SystemChecker struct {
	logger *log.CommonLogger
}

type JsonResult struct { 
    command string
    content string
    error_message string
	os string
	result string
}

type Command struct {
    key string
    command string
	desc string	
}


func (s SystemChecker) Execute(debug_client int, debug_addr string) (bool, string) {

	s.logger = log.GetLogger()
	commander := command.Commander{}


	commands := []Command{
        {"whoami", "whoami", "Running Account"},
		{"os", "cat /etc/*release", "OS (cat /etc/*release)"},
		{"cpu", 
			"grep 'processor' /proc/cpuinfo | tail -1 && grep 'cpu cores' /proc/cpuinfo | tail -1", 
			"CPU/Core"},
		{"memory", "free -h", "Memory (Free -h)"},
		{"disk", "df -h", "DISK(df -h)"},
		{"semaphore",
		"cat /proc/sys/kernel/sem && cat /proc/sys/vm/max_map_count",
		"System Parameter (kerner.sam, vm.max_map_count)"},
		{"ulimit", "ulimit -a", "Ulimit Configuration"},
		{"rpm", "rpm -qa", "RPM List Checking"},
    }

	s.logger.BasicPrint("\n")
	for _, command := range commands {

		s.logger.BasicPrint("|| " + command.desc + "  --------------------")
		
		flag, result := commander.Execute(command.command, debug_client, debug_addr)

		if flag == true{
			s.parseResult(command.key, result)
		}
	}

	return true, ""
}

func (s SystemChecker) parseResult(key string, result string){

	var parseVal map[string]interface{}
	if err := json.Unmarshal([]byte(result), &parseVal); err != nil {
		panic(err)
	}

	content := s.valueFromKey(parseVal, "content")

	if key == "rpm" {
		s.checkRPMList(content)
	} else{
		s.logger.BasicPrint(content)
	}

	s.logger.BasicPrint("\n")
}


func (s SystemChecker) valueFromKey(r map[string]interface{}, key ...string) string {
	if len(key) == 1 {
		return r[key[0]].(string)
	} else if len(key) == 2 {
		return r[key[0]].(map[string]interface{})[key[1]].(string)
	} else {
		fmt.Errorf("Not Support for 3 or more parameters.")
	}
	return ""
}


func (s SystemChecker) checkRPMList(result string) bool{

	isFind := false
	findString := "[OK] "
	notfindString := "[NOT EXIST] "

	RPMList := []string{
		"glibc", "libX11", "libxcb", "libXcomposite", "libXcursor", 
		"libXdamage", "libXext", "libXfixes", "libXi", "libXrender", 
		"libXtst", "glib2", "glib2", "nss", "nss-util", "nss", "nspr", 
		"cups-libs", "dbus-libs", "expat", "libXScrnSaver", "libXrandr", 
		"glib2", "alsa-lib", "pango", "cairo", "at-spi2-atk", "gtk3", 
		"gdk-pixbuf2", "libgcc", "ksh", "libnsl", "GConf2", 
		"libdrm", "mesa-libgbm", "libxshmfence",
	}


	// result parsing
	resultLines := strings.Split(result, "\\n")
	var resultArr []string
	for _, line := range resultLines {
		resultArr = append(resultArr, line)
	}

	for _, rpm := range RPMList {

		isFind = false

		for _, result := range resultArr {
			
			// compare
			if strings.Contains(result, rpm){
				fmt.Println(findString + rpm)
				isFind = true
			}
		}

		if isFind == false{
			fmt.Println(notfindString + rpm)
		}
	}

	return true
}
