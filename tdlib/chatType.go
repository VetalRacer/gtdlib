// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatType Describes the type of a chat
type ChatType interface {
	GetChatTypeEnum() ChatTypeEnum
}

// ChatTypeEnum Alias for abstract ChatType 'Sub-Classes', used as constant-enum here
type ChatTypeEnum string

// ChatType enums
const (
	ChatTypePrivateType    ChatTypeEnum = "chatTypePrivate"
	ChatTypeBasicGroupType ChatTypeEnum = "chatTypeBasicGroup"
	ChatTypeSupergroupType ChatTypeEnum = "chatTypeSupergroup"
	ChatTypeSecretType     ChatTypeEnum = "chatTypeSecret"
)

func unmarshalChatType(rawMsg *json.RawMessage) (ChatType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatTypeEnum(objMap["@type"].(string)) {
	case ChatTypePrivateType:
		var chatTypePrivate ChatTypePrivate
		err := json.Unmarshal(*rawMsg, &chatTypePrivate)
		return &chatTypePrivate, err

	case ChatTypeBasicGroupType:
		var chatTypeBasicGroup ChatTypeBasicGroup
		err := json.Unmarshal(*rawMsg, &chatTypeBasicGroup)
		return &chatTypeBasicGroup, err

	case ChatTypeSupergroupType:
		var chatTypeSupergroup ChatTypeSupergroup
		err := json.Unmarshal(*rawMsg, &chatTypeSupergroup)
		return &chatTypeSupergroup, err

	case ChatTypeSecretType:
		var chatTypeSecret ChatTypeSecret
		err := json.Unmarshal(*rawMsg, &chatTypeSecret)
		return &chatTypeSecret, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ChatTypePrivate An ordinary chat with a user
type ChatTypePrivate struct {
	tdCommon
	UserID int32 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of ChatTypePrivate
func (chatTypePrivate *ChatTypePrivate) MessageType() string {
	return "chatTypePrivate"
}

// NewChatTypePrivate creates a new ChatTypePrivate
//
// @param userID User identifier
func NewChatTypePrivate(userID int32) *ChatTypePrivate {
	chatTypePrivateTemp := ChatTypePrivate{
		tdCommon: tdCommon{Type: "chatTypePrivate"},
		UserID:   userID,
	}

	return &chatTypePrivateTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypePrivate *ChatTypePrivate) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypePrivateType
}

// ChatTypeBasicGroup A basic group (i.e., a chat with 0-200 other users)
type ChatTypeBasicGroup struct {
	tdCommon
	BasicGroupID int32 `json:"basic_group_id"` // Basic group identifier
}

// MessageType return the string telegram-type of ChatTypeBasicGroup
func (chatTypeBasicGroup *ChatTypeBasicGroup) MessageType() string {
	return "chatTypeBasicGroup"
}

// NewChatTypeBasicGroup creates a new ChatTypeBasicGroup
//
// @param basicGroupID Basic group identifier
func NewChatTypeBasicGroup(basicGroupID int32) *ChatTypeBasicGroup {
	chatTypeBasicGroupTemp := ChatTypeBasicGroup{
		tdCommon:     tdCommon{Type: "chatTypeBasicGroup"},
		BasicGroupID: basicGroupID,
	}

	return &chatTypeBasicGroupTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeBasicGroup *ChatTypeBasicGroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeBasicGroupType
}

// ChatTypeSupergroup A supergroup (i.e. a chat with up to GetOption("supergroup_max_size") other users), or channel (with unlimited members)
type ChatTypeSupergroup struct {
	tdCommon
	SupergroupID int32 `json:"supergroup_id"` // Supergroup or channel identifier
	IsChannel    bool  `json:"is_channel"`    // True, if the supergroup is a channel
}

// MessageType return the string telegram-type of ChatTypeSupergroup
func (chatTypeSupergroup *ChatTypeSupergroup) MessageType() string {
	return "chatTypeSupergroup"
}

// NewChatTypeSupergroup creates a new ChatTypeSupergroup
//
// @param supergroupID Supergroup or channel identifier
// @param isChannel True, if the supergroup is a channel
func NewChatTypeSupergroup(supergroupID int32, isChannel bool) *ChatTypeSupergroup {
	chatTypeSupergroupTemp := ChatTypeSupergroup{
		tdCommon:     tdCommon{Type: "chatTypeSupergroup"},
		SupergroupID: supergroupID,
		IsChannel:    isChannel,
	}

	return &chatTypeSupergroupTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSupergroup *ChatTypeSupergroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSupergroupType
}

// ChatTypeSecret A secret chat with a user
type ChatTypeSecret struct {
	tdCommon
	SecretChatID int32 `json:"secret_chat_id"` // Secret chat identifier
	UserID       int32 `json:"user_id"`        // User identifier of the secret chat peer
}

// MessageType return the string telegram-type of ChatTypeSecret
func (chatTypeSecret *ChatTypeSecret) MessageType() string {
	return "chatTypeSecret"
}

// NewChatTypeSecret creates a new ChatTypeSecret
//
// @param secretChatID Secret chat identifier
// @param userID User identifier of the secret chat peer
func NewChatTypeSecret(secretChatID int32, userID int32) *ChatTypeSecret {
	chatTypeSecretTemp := ChatTypeSecret{
		tdCommon:     tdCommon{Type: "chatTypeSecret"},
		SecretChatID: secretChatID,
		UserID:       userID,
	}

	return &chatTypeSecretTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSecret *ChatTypeSecret) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSecretType
}
