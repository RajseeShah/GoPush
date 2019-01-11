package apns_functions

import (
	"crypto/tls"
	"goPush/config"
	"net/http"
	"golang.org/x/net/http2"
	"encoding/json"
)

type ApnsClient struct {
	deviceToken string
	pemFilePath string
	apnsHost string
	httpClient *http.Client
}

const maxPayload  = 4096 // 4KB at most


func NewApnsClient(deviceToken, pemFilePath string, apnsMode int) (*ApnsClient, error) {

	apnsHost := config.ApnsSendBoxEndPointUrl

	if apnsMode == 2{
		apnsHost = config.ApnsProductionEndPointUrl
	}

	certificate, err := tls.LoadX509KeyPair(pemFilePath, pemFilePath)

	if err!= nil{
		return nil, err
	}

	confingurations := &tls.Config{
		Certificates : []tls.Certificate{certificate},
	}

	confingurations.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig:confingurations}

	err = http2.ConfigureTransport(transport)
	if err!=nil{
		return nil, err
	}

	apnsClient := &ApnsClient{

			deviceToken:deviceToken,
			pemFilePath:pemFilePath,
			apnsHost:apnsHost,
			httpClient:&http.Client{Transport:transport},
	}

	return apnsClient, nil
}

func (client *ApnsClient) SendApnsPush(header *ApnsHeader, payload *ApnsPayload) (*Response, error){

	apnsPayload, err := json.Marshal(payload)
	if err!=nil{
		return nil, err
	}

	// check payload length before even hitting Apple.
	if len(apnsPayload) > maxPayload {
		return &Response{
			Code:400,
			Message:"Payload too large",
		}, nil
	}

	//Sending request to APNS server
	finalResponse, err := client.sendRequestToApnsServer(header, apnsPayload)

	if err!= nil{
		finalResponse.Code = 400
		finalResponse.Message = err.Error()
	}

	return finalResponse, nil


}



