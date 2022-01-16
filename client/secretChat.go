// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline request
// @param secretChatID Secret chat identifier
func (client *Client) GetSecretChat(secretChatID int32) (*tdlib.SecretChat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getSecretChat",
		"secret_chat_id": secretChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var secretChatDummy tdlib.SecretChat
	err = json.Unmarshal(result.Raw, &secretChatDummy)
	return &secretChatDummy, err

}
