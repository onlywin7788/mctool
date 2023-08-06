package command

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"runtime"
	log "bits/modules/common/log"
)

type Commander struct {
}

func (commander *Commander) Execute(cmd string, debug int) (bool, string) {
	
	var result bool
	var returnString string

	if debug == 1 {
		result,returnString = CallClient(cmd)
	} else{
		result,returnString = SendCommand(cmd)
	}

	return result, returnString
}

func SendCommand(cmd string) (bool, string) {

	var returnVal = true
	var returnError = ""

	var cmdVal1, cmdVal2 = "sh", "-c"
	if runtime.GOOS == "windows" {
		cmdVal1, cmdVal2 = "cmd", "/C"
	}

	output, err := exec.Command(cmdVal1, cmdVal2, cmd).Output()
	
	if err != nil {
		returnVal = false
		returnError = err.Error()
	}

	retunVal := string(output[:])

	//convert json
	data := make(map[string]interface{})
	data["result"] = "successed"
	data["error_message"] = returnError
	data["os"] = runtime.GOOS
	data["command"] = cmd
	data["content"] = retunVal
	doc, _ := json.Marshal(data)
	retunJson := string(doc[:])

	return returnVal, retunJson
}


func (commander Commander) DebugServerListen() {
	
	logger := log.CommonLogger{}
	logger.Trace("Remote REST Debugger Server : 34972")

	http.HandleFunc("/debug", func(rw http.ResponseWriter, r *http.Request) {
		cmd := r.Header.Get("command")

		_, output := SendCommand(cmd)

		rw.Write([]byte(output))
    })
	http.ListenAndServe(":34972", nil)
}

func CallClient(cmd string) (bool, string){
	logger := log.CommonLogger{}
	logger.Dummy()

    req, err := http.NewRequest("GET", "http://10.10.19.157:34972/debug", nil)
    if err != nil {
        panic(err)
    }
 
    req.Header.Add("command", cmd)
 
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
 
    bytes, _ := ioutil.ReadAll(resp.Body)
    str := string(bytes)
    
	return true, str
}