package apns_functions

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

//Send push to Ios device using p8/pem file
func (client *ApnsClient) sendRequestToApnsServer(header *ApnsHeader, data []byte) (*Response, error) {

	requestUrl := fmt.Sprintf("%v/3/device/%v", client.apnsHost, client.deviceToken)

	fmt.Println("--------Sending request to APNS server-----------")

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("--------Error in request-----------")
		return nil, err
	}

	request.Header.Add("Content-type", "application/json")

	//Set authorization for sending push using p8 file
	if client.fileType == 2{

		request.Header.Add("authorization", "bearer "+client.jwtToken)
	}

	header.set(request.Header)

	response, err := client.httpClient.Do(request)

	if err != nil {
		fmt.Println("--------Error in response-----------")
		return nil, err
	}

	defer response.Body.Close()

	finalResponse := new(Response)
	finalResponse.Code = response.StatusCode
	finalResponse.ApnsId = response.Header.Get("apns-id")

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("--------Error while reading response-----------")
		finalResponse.Message = err.Error()
		return finalResponse, err
	}

	res := APNSResponse{}

	//If error occured then and then we'll get something in body
	if len(responseBody) > 0 {

		err = json.Unmarshal(responseBody, &res)

		if err != nil {
			fmt.Println("--------Error while unmarshling response---------")
			finalResponse.Message = err.Error()
			return finalResponse, nil
		}

		if res.Reason != ""{
			finalResponse.Message = "Push notification is failed"
			finalResponse.Error = res
		}
	}

	if res.Reason == ""{
		finalResponse.Message = "Push Notification sent successfully."
	}

	return finalResponse, nil
}


// set headers for an HTTP request
func (h *ApnsHeader) set(reqHeader http.Header) {
	// headers are optional
	if h == nil {
		return
	}

	if h.ApnsId != "" {
		reqHeader.Set("apns-id", h.ApnsId)
	} // when omitted, Apple will generate a UUID for you

	if h.ApnsCollapseId != "" {
		reqHeader.Set("apns-collapse-id", h.ApnsCollapseId)
	}

	if h.ApnsPriority != "" {
		reqHeader.Set("apns-priority", h.ApnsPriority)
	} // when omitted, the default priority is 10

	if h.ApnsTopic != "" {
		reqHeader.Set("apns-topic", h.ApnsTopic)
	}
	if !h.ApnsExpiration.IsZero() {
		reqHeader.Set("apns-expiration", fmt.Sprintf("%v", h.ApnsExpiration.Unix()))
	}

}
