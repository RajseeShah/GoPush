package fcm_functions

import (
	"fmt"
	"goPush/config"
	"strings"
)

type FcmPayLoadData struct {
	Token                    string                 `json:"to,omitempty"`
	RegistrationIDs          []string               `json:"registration_ids,omitempty"`
	Condition                string                 `json:"condition,omitempty"`
	CollapseKey              string                 `json:"collapse_key,omitempty"`
	Priority                 string                 `json:"priority,omitempty"`
	ContentAvailable         bool                   `json:"content_available,omitempty"`
	DelayWhileIdle           bool                   `json:"delay_while_idle,omitempty"`
	TimeToLive               int                    `json:"time_to_live,omitempty"`
	DeliveryReceiptRequested bool                   `json:"delivery_receipt_requested,omitempty"`
	DryRun                   bool                   `json:"dry_run,omitempty"`
	Notification             Notification          `json:"notification,omitempty"`
	Data                     map[string]interface{} `json:"data,omitempty"`
}

type Notification struct {
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Sound        string `json:"sound,omitempty"`
	Badge        string `json:"badge,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Color        string `json:"color,omitempty"`
	ClickAction  string `json:"click_action,omitempty"`
	BodyLocKey   string `json:"body_loc_key,omitempty"`
	BodyLocArgs  string `json:"body_loc_args,omitempty"`
	TitleLocKey  string `json:"title_loc_key,omitempty"`
	TitleLocArgs string `json:"title_loc_args,omitempty"`
}

//This function will validate FCM payload data set by User
func (data *FcmPayLoadData) validate()  string {

	fmt.Println("------------Validating Data-----------------")
	errorMessage := ""

	if data == nil {
		return config.PayloadInvalid
	}

	// validate target identifier: `to` or `condition`, or `registration_ids`
	opCnt := strings.Count(data.Condition, "&&") + strings.Count(data.Condition, "||")
	if data.Token == "" && (data.Condition == "" || opCnt > 2) && len(data.RegistrationIDs) == 0 {
		return config.TokenInvalid
	}
	if len(data.RegistrationIDs) > 1000 {
		return config.TooManyIds
	}
	if data.TimeToLive > 2419200 {
		return config.TimeToLiveInvalid
	}

	fmt.Println("------------Validated successfully-----------------")

	return errorMessage
}
