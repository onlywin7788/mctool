package util

import (
	"fmt"
	"strings"
	log "bits/modules/common/log"
	command "bits/modules/command"
	"encoding/json"
)

type SystemChecker struct {
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


func (systemChecker SystemChecker) Check(debug_client int) (bool, string) {

	logger := log.CommonLogger{}
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

	logger.BasicPrint("\n")
	for _, command := range commands {

		logger.BasicPrint("|| " + command.desc + "  --------------------")
		
		flag, result := commander.Execute(command.command, debug_client)

		if flag == true{
			ParseResult(command.key, result)
		}
	}

	return true, ""
}

func ParseResult(key string, result string){
	
	logger := log.CommonLogger{}
	logger.Dummy()

	var parseVal map[string]interface{}
	if err := json.Unmarshal([]byte(result), &parseVal); err != nil {
		panic(err)
	}

	content := valueFromKey(parseVal, "content")

	if key == "rpm" {
		checkRPMList(content)
	} else{
		logger.BasicPrint(content)
	}

	logger.BasicPrint("\n")
}


func valueFromKey(r map[string]interface{}, key ...string) string {
	if len(key) == 1 {
		return r[key[0]].(string)
	} else if len(key) == 2 {
		return r[key[0]].(map[string]interface{})[key[1]].(string)
	} else {
		fmt.Errorf("Not Support for 3 or more parameters.")
	}
	return ""
}


func checkRPMList(result string) bool{

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
