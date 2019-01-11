package apns_functions

type Response struct {
	Code    int         `json:"code"`
	Error    APNSResponse `json:"error,omitempty"`
	ApnsId string `json:"apns_id"`
	Message    string `json:"message"`
}

type APNSResponse struct {
	Reason string `json:"reason,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}
