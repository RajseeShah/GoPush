package fcm_functions

type Response struct {
	Code    int         `json:"code"`
	Data    FcmResponse `json:"data,omitempty"`
	Message    string `json:"message"`
}

type FcmResponse struct {
	MulticastId int `json:"multicast_id,omitempty"`
	Success int `json:"success,omitempty"`
	Failure int `json:"failure,omitempty"`
	CanonicalId int `json:"canonical_id,omitempty"`
	Results []FcmResult `json:"results,omitempty"`
}

type FcmResult struct {
	MessageId interface{} `json:"message_id,omitempty"`
	RegistrationId interface{} `json:"registration_id,omitempty"`
	Error string `json:"error,omitempty"`
}

