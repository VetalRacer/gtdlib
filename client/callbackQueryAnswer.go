// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param chatID Identifier of the chat with the message
// @param messageID Identifier of the message from which the query originated
// @param payload Query payload
func (client *Client) GetCallbackQueryAnswer(chatID int64, messageID int64, payload tdlib.CallbackQueryPayload) (*tdlib.CallbackQueryAnswer, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatID,
		"message_id": messageID,
		"payload":    payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callbackQueryAnswer tdlib.CallbackQueryAnswer
	err = json.Unmarshal(result.Raw, &callbackQueryAnswer)
	return &callbackQueryAnswer, err

}
