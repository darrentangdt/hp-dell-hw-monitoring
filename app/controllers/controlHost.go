package controllers

import (
	"bytes"
	"fmt"
	"tms/app/models"
)

func ControlHost(monitor models.Monitor, body string) {
	fmt.Println(body)
	var url string
	if monitor.Manuf == models.DELL {
		url = "https://" + monitor.Host + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
		if body == "on" {
			body = `{"ResetType":"On"}`
		} else if body == "off" {
			body = `{"ResetType":"ForceOff"}`
		} else if body == "reset" {
			body = `{"ResetType":"GracefulRestart"}`
		} else {
		}
	} else {
		url = "https://" + monitor.Host + "/rest/v1/Systems/1"
		//url = "https://httpbin.org/post"
		if body == "off" {
			body = `{"Action":"Reset","ResetType":"ForceOff"}`
		} else if body == "on" {
			body = `{"Action":"Reset","ResetType":"On"}`
		} else if body == "reset" {
			body = `{"Action":"Reset","ResetType":"ForceRestart"}`
		} else {
		}
	}
	reqBody := bytes.NewBufferString(body)
	HttpPostRequest(url, monitor, reqBody)
}
