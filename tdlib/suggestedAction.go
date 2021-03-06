// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// SuggestedAction Describes an action suggested to the current user
type SuggestedAction interface {
	GetSuggestedActionEnum() SuggestedActionEnum
}

// SuggestedActionEnum Alias for abstract SuggestedAction 'Sub-Classes', used as constant-enum here
type SuggestedActionEnum string

// SuggestedAction enums
const (
	SuggestedActionEnableArchiveAndMuteNewChatsType SuggestedActionEnum = "suggestedActionEnableArchiveAndMuteNewChats"
	SuggestedActionCheckPhoneNumberType             SuggestedActionEnum = "suggestedActionCheckPhoneNumber"
)

func unmarshalSuggestedAction(rawMsg *json.RawMessage) (SuggestedAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SuggestedActionEnum(objMap["@type"].(string)) {
	case SuggestedActionEnableArchiveAndMuteNewChatsType:
		var suggestedActionEnableArchiveAndMuteNewChats SuggestedActionEnableArchiveAndMuteNewChats
		err := json.Unmarshal(*rawMsg, &suggestedActionEnableArchiveAndMuteNewChats)
		return &suggestedActionEnableArchiveAndMuteNewChats, err

	case SuggestedActionCheckPhoneNumberType:
		var suggestedActionCheckPhoneNumber SuggestedActionCheckPhoneNumber
		err := json.Unmarshal(*rawMsg, &suggestedActionCheckPhoneNumber)
		return &suggestedActionCheckPhoneNumber, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// SuggestedActionEnableArchiveAndMuteNewChats Suggests the user to enable "archive_and_mute_new_chats_from_unknown_users" option
type SuggestedActionEnableArchiveAndMuteNewChats struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionEnableArchiveAndMuteNewChats
func (suggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats) MessageType() string {
	return "suggestedActionEnableArchiveAndMuteNewChats"
}

// NewSuggestedActionEnableArchiveAndMuteNewChats creates a new SuggestedActionEnableArchiveAndMuteNewChats
//
func NewSuggestedActionEnableArchiveAndMuteNewChats() *SuggestedActionEnableArchiveAndMuteNewChats {
	suggestedActionEnableArchiveAndMuteNewChatsTemp := SuggestedActionEnableArchiveAndMuteNewChats{
		tdCommon: tdCommon{Type: "suggestedActionEnableArchiveAndMuteNewChats"},
	}

	return &suggestedActionEnableArchiveAndMuteNewChatsTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionEnableArchiveAndMuteNewChats *SuggestedActionEnableArchiveAndMuteNewChats) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionEnableArchiveAndMuteNewChatsType
}

// SuggestedActionCheckPhoneNumber Suggests the user to check authorization phone number and change the phone number if it is inaccessible
type SuggestedActionCheckPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of SuggestedActionCheckPhoneNumber
func (suggestedActionCheckPhoneNumber *SuggestedActionCheckPhoneNumber) MessageType() string {
	return "suggestedActionCheckPhoneNumber"
}

// NewSuggestedActionCheckPhoneNumber creates a new SuggestedActionCheckPhoneNumber
//
func NewSuggestedActionCheckPhoneNumber() *SuggestedActionCheckPhoneNumber {
	suggestedActionCheckPhoneNumberTemp := SuggestedActionCheckPhoneNumber{
		tdCommon: tdCommon{Type: "suggestedActionCheckPhoneNumber"},
	}

	return &suggestedActionCheckPhoneNumberTemp
}

// GetSuggestedActionEnum return the enum type of this object
func (suggestedActionCheckPhoneNumber *SuggestedActionCheckPhoneNumber) GetSuggestedActionEnum() SuggestedActionEnum {
	return SuggestedActionCheckPhoneNumberType
}
