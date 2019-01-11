package apns_functions

import "time"

type ApnsHeader struct {
	Authorization 		string	 	`json:"authorization,omitempty"`
	ApnsId         		string      `json:"apns-id,omitempty"`
	ApnsExpiration 		time.Time    `json:"apns-expiration,omitempty"`
	ApnsPriority   		string      `json:"apns-priority,omitempty"`
	ApnsTopic     	 	string    	`json:"apns-topic,omitempty"`
	ApnsCollapseId 		string      `json:"apns-collapse-id,omitempty"`
}

type ApnsPayload struct {
	Aps struct {
		Alert            interface{} `json:"alert,omitempty"`
		Badge            int         `json:"badge,omitempty"`
		Sound            string      `json:"sound,omitempty"`
		ThreadId         string      `json:"thread-id,omitempty"`
		Category         string      `json:"category,omitempty"`
		ContentAvailable int         `json:"content-available,omitempty"`
		MutableContent   int         `json:"mutable-content,omitempty"`
	} `json:"aps"`
}

type Alert struct {
	Title 			string				`json:"title,omitempty"`
	Subtitle 		string				`json:"subtitle,omitempty"`
	Body 	   		string				`json:"body,omitempty"`
	LaunchImage		string				`json:"launch-image,omitempty"`
	TitleLocKey 	string			    `json:"title-loc-key,omitempty"`
	TitleLocArgs 	[]string			`json:"title-loc-args,omitempty"`
	SubtitleLocKey 	string				`json:"subtitle-loc-key,omitempty"`
	SubtitleLocArgs []string			`json:"subtitle-loc-args,omitempty"`
	LocKey 			string				`json:"loc-key,omitempty"`
	LocArgs 		[]string			`json:"loc-args,omitempty"`
}



