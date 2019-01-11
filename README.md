# GoPush

#Your main.go File

package main

import (
	"fmt"
	"encoding/json"
	"goPush/apns_functions"
	"time"
	"goPush/fcm_functions"
)

const fcmServerKey  = "AIzaSyAm4UiXoCus97g5xZoM8rsVAjrvRBs_lDU"  //Your FCM server key
const pemFilePath  = "/var/www/html/pushcert.pem" //Pem file path
const apnsMode = 2  //1 -Sandbox , 2-Production

func main()  {

	fmt.Println("---------------------------------APNS PUSH START--------------------------------------")

	//Device token of IOS device
	deviceToken := "63B649B2D61D1DE2333C2FFB05450AE723A8FC0DF19E5261F89B051A0B8ABFDA"


	//Create APNS client
	/*  PARAM          		DESCRIPTION                              Type
		deviceToken		 - Device token of IOS device		    	- string
		pemFilePath		 - path Where your PEM file resides 		- string
		apnsMode   		 - 1/2 (1 -Sandbox , 2-Production)			- int
	 */
	apnsClient,_ := apns_functions.NewApnsClient(deviceToken, pemFilePath, apnsMode)


	//Prepare APNS payload data
	apnsPayload := apns_functions.ApnsPayload{}


	// Set Alert part for APNS push, there are many params you can set for alert, I have used only three
	// For additional param detail, check Alert struct
	apnsPayload.Aps.Alert = apns_functions.Alert{
		Title:"YOUR_MESSAGE_TITLE",
		Body:"YOUR_MESSAGE_BODY",
		Subtitle:"YOUR_MESSAGE_SUBTITLE",
	}

	//Set payload's additional params, check Aps struct for addidtional params
	apnsPayload.Aps.Sound = "default"
	apnsPayload.Aps.Badge = 1


	//Set APNS header, only ApnsTopic is mandatory
	apnsHeader := apns_functions.ApnsHeader{
		ApnsPriority:"10",
		ApnsTopic:"com.slideshowmaker.photoeditor", //Mandatory
		ApnsExpiration:time.Now(),
	}


	//Call this function for sending push notification!
	resultAPNS,_ := apnsClient.SendApnsPush(&apnsHeader,&apnsPayload)

	//Marshal the result in order to get response as JSON
	responseAPNS, _ := json.Marshal(resultAPNS)
	fmt.Println(string(responseAPNS))


	fmt.Println("---------------------------------APNS PUSH END--------------------------------------")



	fmt.Println("-----------------------------------FCM PUSH START------------------------------------")

	//Create FCM client first by passing FCM server key
	fcmClient := fcm_functions.NewFcmClient(fcmServerKey)

	//Prepare FCM payload data
	data := new(fcm_functions.FcmPayLoadData)

	//If you want to send FCM Push notification to IOS device
	data.Token = "dctYDZiwcjk:APA91bHjFPjxe0Dtw87v7YzYkhJMQL1zGoFCS5aPBgKQvj1t0KKE_fsXFqnTsklCRqHlkFaARHGlQNCVtNYA0VMrb7wRuthkwlmNGZXWYU23g9nvwuCbbmj4hSIOk_cYKA0i77DV4NsH"

	//If you want to send FCM Push notification to android devices
	registrationIds := []string{"your_registration_id_1","your_registration_id_2","your_registration_id_3"}
	data.RegistrationIDs =registrationIds

	//Comman Parameters for any push (IOS/Android)
	data.Priority = "High"
	data.Notification.Title = "Test message"
	data.Notification.Body = "Test message"

	// You can also set additional parameters of FcmPayloadData according to requirements
	// like,
		data.Data =  map[string]interface{}{
       		"key": "value",
     		}


	//Call function for sending push to IOS/Android device using FCM client reference and pass FCM payload data to it
	result, _ := fcmClient.SendFcmPushNotification(data)

	//Marshal the result in order to get response as JSON
	responseFCM, _ := json.Marshal(result)
	fmt.Println(string(responseFCM))

	fmt.Println("------------------------------------FCM PUSH END--------------------------------------")
}


