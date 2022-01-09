// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles
// @param chatID Chat identifier
func (client *Client) GetChatAdministrators(chatID int64) (*tdlib.ChatAdministrators, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChatAdministrators",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatAdministrators tdlib.ChatAdministrators
	err = json.Unmarshal(result.Raw, &chatAdministrators)
	return &chatAdministrators, err

}
