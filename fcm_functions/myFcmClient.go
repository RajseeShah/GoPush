package fcm_functions

import (
	"net/http"
	"encoding/json"
)

type FcmClient struct {
	fcmServerKey string
	httpClient *http.Client
}

//This function will create new HTTP client for sending http request to FCM server
func NewFcmClient(fcmServerKey string) (*FcmClient){

	fcmClient := &FcmClient{

		fcmServerKey : fcmServerKey,
		httpClient : &http.Client{},

	}

	return fcmClient
}

//This function will send push request to FCM server
func (client *FcmClient) SendFcmPushNotification(data *FcmPayLoadData) (*Response, error)  {

	if errorMessage := data.validate(); errorMessage != "" {

		finalResponse := new(Response)
		finalResponse.Code = 400
		finalResponse.Message = errorMessage

		return finalResponse, nil
	}

	// marshal data
	fcmPayLoadData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	//Sending request to FCM server
	finalResponse, err := client.sendRequestToFCMServer(fcmPayLoadData)

	if err!= nil{
		finalResponse.Code = 400
		finalResponse.Message = err.Error()
	}

	return finalResponse, nil
}



