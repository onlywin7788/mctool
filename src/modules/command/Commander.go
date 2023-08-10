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
	logger *log.CommonLogger
}

func (c Commander) Execute(cmd string, debug int, debug_addr string) (bool, string) {

	var result bool
	var returnString string
	c.logger = log.GetLogger()

	if debug == 1 {
		result,returnString = c.callClient(cmd, debug_addr)
	} else{
		result,returnString = c.sendCommand(cmd)
	}

	return result, returnString
}

func (c Commander) DebugServerListen(port string) {
	
	c.logger = log.GetLogger()

	c.logger.Info("Remote REST Debugger Server Listen : " + port)

	http.HandleFunc("/debug", func(rw http.ResponseWriter, r *http.Request) {
		cmd := r.Header.Get("command")

		_, output := c.sendCommand(cmd)

		rw.Write([]byte(output))
    })
	http.ListenAndServe(":" + port, nil)
}

func (c Commander) sendCommand(cmd string) (bool, string) {

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

func (c Commander)callClient(cmd string, debug_addr string) (bool, string){
	
    req, err := http.NewRequest("GET", "http://" + debug_addr + "/debug", nil)
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