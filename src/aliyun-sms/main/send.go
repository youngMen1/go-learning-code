package main

import (
	"encoding/json"
	"fmt"
	"github.com/KenmyZhang/aliyun-communicate"
)

var (
	gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
	accessKeyId     = "*********"
	accessKeySecret = "***********"
	phoneNumbers    = "18682333447"
	signName        = "古德菲力健身"
	templateCode    = "SMS_161380087"
	templateParam   = "{\"code\":\"123456\"}"
)

func main() {
	fmt.Print("111111111111")
	smsClient := aliyunsmsclient.New(gatewayUrl)
	fmt.Print("111111111111")
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	fmt.Println("Got raw response from server:", string(result.RawResponse))
	if err != nil {
		panic("Failed to send Message: " + err.Error())
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	if result.IsSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
	} else {
		fmt.Println("Failed to send a SMS:", resultJson)
	}
}
