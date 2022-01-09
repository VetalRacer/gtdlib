// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// CreateChatFilter Creates new chat filter. Returns information about the created chat filter
// @param filter Chat filter
func (client *Client) CreateChatFilter(filter *tdlib.ChatFilter) (*tdlib.ChatFilterInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":  "createChatFilter",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo tdlib.ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}

// EditChatFilter Edits existing chat filter. Returns information about the edited chat filter
// @param chatFilterID Chat filter identifier
// @param filter The edited chat filter
func (client *Client) EditChatFilter(chatFilterID int32, filter *tdlib.ChatFilter) (*tdlib.ChatFilterInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "editChatFilter",
		"chat_filter_id": chatFilterID,
		"filter":         filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo tdlib.ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}
