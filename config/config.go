package config

const (

	// endpoint URL of FCM service.
	FcmEndPointUrl = "https://fcm.googleapis.com/fcm/send"

	//FCM payload data validation errors
	PayloadInvalid = "Payload data is invalid"
	TokenInvalid = "Device token is invalid or registration ids are not set"
	TooManyIds = "Too many registrations ids"
	TimeToLiveInvalid = "Time to live is invalid"

	// Endpoint URLs for APNS
	ApnsSendBoxEndPointUrl = "https://api.development.push.apple.com"
	ApnsProductionEndPointUrl = "https://api.push.apple.com"

)
