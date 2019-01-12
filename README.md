<h1><a id="GoPush_0"></a>GoPush</h1>
<p>#Your main.go File</p>
<p>package main</p>
<p>import (<br>
“fmt”<br>
“encoding/json”<br>
“goPush/apns_functions”<br>
“time”<br>
“goPush/fcm_functions”<br>
)</p>
<p>const fcmServerKey  = “your_fcm_server_key”  //Your FCM server key<br>
const pemFilePath  = “your_pem_file_path” //Pem file path<br>
const apnsMode = 2  //1 -Sandbox , 2-Production</p>
<p>func main()  {</p>
<pre><code>fmt.Println(&quot;---------------------------------APNS PUSH START--------------------------------------&quot;)

//Device token of IOS device
deviceToken := &quot;YOUR_IOS_DEVICE_TOKEN&quot;


//Create APNS client
/*  PARAM               DESCRIPTION                              Type
    deviceToken      - Device token of IOS device               - string
    pemFilePath      - path Where your PEM file resides         - string
    apnsMode         - 1/2 (1 -Sandbox , 2-Production)          - int
 */
apnsClient,_ := apns_functions.NewApnsClient(deviceToken, pemFilePath, apnsMode)


//Prepare APNS payload data
apnsPayload := apns_functions.ApnsPayload{}


// Set Alert part for APNS push, there are many params you can set for alert, I have used only three
// For additional param detail, check Alert struct
apnsPayload.Aps.Alert = apns_functions.Alert{
    Title:&quot;YOUR_MESSAGE_TITLE&quot;,
    Body:&quot;YOUR_MESSAGE_BODY&quot;,
    Subtitle:&quot;YOUR_MESSAGE_SUBTITLE&quot;,
}

//Set payload's additional params, check Aps struct for addidtional params
apnsPayload.Aps.Sound = &quot;default&quot;
apnsPayload.Aps.Badge = 1


//Set APNS header, only ApnsTopic is mandatory
apnsHeader := apns_functions.ApnsHeader{
    ApnsPriority:&quot;10&quot;,
    ApnsTopic:&your_apns_topic&quot;, //Mandatory
    ApnsExpiration:time.Now(),
}


//Call this function for sending push notification!
resultAPNS,_ := apnsClient.SendApnsPush(&amp;apnsHeader,&amp;apnsPayload)

//Marshal the result in order to get response as JSON
responseAPNS, _ := json.Marshal(resultAPNS)
fmt.Println(string(responseAPNS))


fmt.Println(&quot;---------------------------------APNS PUSH END--------------------------------------&quot;)



fmt.Println(&quot;-----------------------------------FCM PUSH START------------------------------------&quot;)

//Create FCM client first by passing FCM server key
fcmClient := fcm_functions.NewFcmClient(fcmServerKey)

//Prepare FCM payload data
data := new(fcm_functions.FcmPayLoadData)

//If you want to send FCM Push notification to IOS device
data.Token = &quot;YOUR_IOS_DEVICE_TOKEN&quot;

//If you want to send FCM Push notification to android devices
registrationIds := []string{&quot;your_registration_id_1&quot;,&quot;your_registration_id_2&quot;,&quot;your_registration_id_3&quot;}
data.RegistrationIDs =registrationIds

//Comman Parameters for any push (IOS/Android)
data.Priority = &quot;High&quot;
data.Notification.Title = &quot;Test message&quot;
data.Notification.Body = &quot;Test message&quot;

// You can also set additional parameters of FcmPayloadData according to requirements
// like,
    data.Data =  map[string]interface{}{
        &quot;key&quot;: &quot;value&quot;,
        }


//Call function for sending push to IOS/Android device using FCM client reference and pass FCM payload data to it
result, _ := fcmClient.SendFcmPushNotification(data)

//Marshal the result in order to get response as JSON
responseFCM, _ := json.Marshal(result)
fmt.Println(string(responseFCM))

fmt.Println(&quot;------------------------------------FCM PUSH END--------------------------------------&quot;)
</code></pre>
<p>}</p>
