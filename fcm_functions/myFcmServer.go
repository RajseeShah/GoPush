package fcm_functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"goPush/config"
	"bytes"
	"encoding/json"
)

//This function will make request to FCM endpoint and send push to device
func (client *FcmClient) sendRequestToFCMServer(data []byte) (*Response, error){

	fmt.Println("--------Sending request to FCM server-----------")

	request, err := http.NewRequest("POST", config.FcmEndPointUrl, bytes.NewBuffer(data))

	if err!= nil{
		fmt.Println("--------Error in request-----------")
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("key=%s", client.fcmServerKey))
	request.Header.Add("Content-type", "application/json")

	response, err := client.httpClient.Do(request)

	if err != nil{
		fmt.Println("--------Error in response-----------")
		return nil, err
	}

	defer response.Body.Close()

	finalResponse := new(Response)
	finalResponse.Code = response.StatusCode

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil{
		fmt.Println("--------Error while reading response-----------")
		finalResponse.Message = err.Error()
		return finalResponse, err
	}

	res := FcmResponse{}

	err = json.Unmarshal(responseBody, &res)

	if err != nil{
		fmt.Println("--------Error while unmarshling response---------")
		finalResponse.Message = "Invalid FCM server key"
		return finalResponse, nil
	}

	for _,v := range res.Results{

		if v.Error != ""{
			finalResponse.Message =  v.Error
			finalResponse.Code = 400
		}
	}

	if finalResponse.Message == ""{
		finalResponse.Message = "Push Notification sent successfully."
	}

	finalResponse.Data = res

	return finalResponse, nil
}
