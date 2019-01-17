package apns_functions

import (
	"crypto/tls"
	"goPush/config"
	"net/http"
	"golang.org/x/net/http2"
	"encoding/json"
	"goPush/authorization_token"
	"fmt"
	"errors"
)

type ApnsClient struct {
	deviceToken string
	filePath string
	apnsHost string
	httpClient *http.Client
	fileType int //1 - PEM file, 2 - P8 file
	jwtToken string
}

const maxPayload  = 4096 // 4KB at most

//APNS client for sending iOS push using .PEM or .p8 file
func NewApnsClient(deviceToken, filePath string, apnsMode int, teamId, keyId string) (*ApnsClient, error) {

	apnsHost := config.ApnsSendBoxEndPointUrl

	if apnsMode == 2{
		apnsHost = config.ApnsProductionEndPointUrl
	}

	//Check file type whether it's pem or p8
	fileType, err := checkFileType(filePath)

	if err!= nil{
		return nil, err
	}

	apnsClient := new(ApnsClient)
	myHttpClient := &http.Client{} //For p8 file, we'll need simple http client
	jwtToken := ""

	//If pem file found, then we've to authenticate it and create httpclient by passing transport
	if fileType == 1{

		certificate, err := tls.LoadX509KeyPair(filePath, filePath)

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

		myHttpClient = &http.Client{Transport:transport} //For pem file, we'll http client with transport

	}else{

		if teamId != "" && keyId != ""{

			//We'll need jwtToken for sending push using p8 file
			jwtToken, err = authorization_token.GenerateAuthorizationToken(filePath, keyId, teamId)

			if err!= nil{
				return nil, err
			}

		}else{

			return nil, errors.New("Key id and/or team id can't be empty.")
		}

	}

	apnsClient = &ApnsClient{

		deviceToken:deviceToken,
		filePath:filePath,
		apnsHost:apnsHost,
		httpClient:myHttpClient,
		fileType:fileType,
		jwtToken:jwtToken,
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

func checkFileType(filePath string) (int, error) {

	fmt.Println("---------------------Checking File type--------------------------")

	fileType := filePath[len(filePath)-3:len(filePath)]

	if fileType == "pem" || fileType == "PEM"{

		fmt.Println("---------------------It's PEM file--------------------------")
		return 1, nil

	}else if fileType == ".p8" || fileType == ".P8"{

		fmt.Println("---------------------It's P8 file--------------------------")
		return 2, nil

	}else{

		fmt.Println("---------------------Not valid file type--------------------------")
		return 0, errors.New("Invalid file type")

	}
}



