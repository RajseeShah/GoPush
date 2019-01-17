<h1><a id="GoPush_0"></a>GoPush</h1>
<p>#Your main.go File</p>
<p>package main</p>
<p>import (<br>
“time”<br>
“fmt”<br>
“os”<br>
“encoding/json”<br>
“goPush/apns_functions”<br>
“goPush/fcm_functions”<br>
)</p>
<p>const fcmServerKey = “your_fcm_server_key”//Your FCM server key for FCM Push</p>
<p>const pemFilePath  = “your_pem_file_path”                                        //Pem file path for APNS push through PEM file<br>
const apnsMode = 2                                                                       //1 -Sandbox , 2-Production</p>
<p>const p8FilePath  = “your_p8_file_path”                 //p8 file path for APNS push through p8 file<br>
const apnsKeyId  = “your_key_id”                                                        //key id for APNS push through p8 file<br>
const teamId  = “your_team_id”                                                          //team id for APNS push through p8 file</p>
<p>func main()  {</p>
<pre><code>fmt.Println(&quot;---------------------------------APNS PUSH START (With PEM/P8 FILE)--------------------------------------&quot;)

//Device token of IOS device
deviceToken := &quot;your_device_token&quot; 


//Create APNS client
/*  PARAM               DESCRIPTION                              Type
    deviceToken      - Device token of IOS device               - string (Mandatory)
    filePath         - path Where your PEM/p8 file resides      - string (Mandatory)
    apnsMode         - 1/2 (1 -Sandbox , 2-Production)          - int (Mandatory)
    teamId           - teamId when using p8 file                - mandaory only when you use p8 file
    apnsKeyId        - apnsKeyId when using p8 file             - mandaory only when you use p8 file
*/
apnsClient,err := apns_functions.NewApnsClient(deviceToken, p8FilePath, apnsMode, teamId, apnsKeyId)

if err!=nil {
    fmt.Println(err)
    os.Exit(0)
}

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
    ApnsTopic: &quot;your_apns_topic&quot;,  //Mandatory
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
data.Token = &quot;your_device_token_of_ios&quot;

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
