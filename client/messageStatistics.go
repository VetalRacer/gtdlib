// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if Message.can_get_statistics == true
// @param chatID Chat identifier
// @param messageID Message identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetMessageStatistics(chatID int64, messageID int64, isDark bool) (*tdlib.MessageStatistics, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getMessageStatistics",
		"chat_id":    chatID,
		"message_id": messageID,
		"is_dark":    isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageStatistics tdlib.MessageStatistics
	err = json.Unmarshal(result.Raw, &messageStatistics)
	return &messageStatistics, err

}
