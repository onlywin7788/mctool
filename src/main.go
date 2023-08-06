package main


import (
	"fmt"
	"os"
	"runtime"
	info "bits/modules/information"
	svc "bits/modules/service"
)

func main(){

	serviceMain := svc.ServiceMain{}

	// if no paramaeter, printing program information 
	if len(os.Args) == 1 {

        fmt.Printf("\n%s (%s) (OS : %s)\n", info.PROGRAM_NAME, info.PROGRAM_DESC, runtime.GOOS)
		fmt.Printf("%s\n\n", info.PROGRAM_COPYRIGHT)
    
	} else {
		serviceMain.Execute()
	}
}