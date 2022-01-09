// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only
// @param chatID The chat that contains the message with the game
// @param messageID Identifier of the message
// @param userID User identifier
func (client *Client) GetGameHighScores(chatID int64, messageID int64, userID int32) (*tdlib.GameHighScores, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getGameHighScores",
		"chat_id":    chatID,
		"message_id": messageID,
		"user_id":    userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores tdlib.GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only
// @param inlineMessageID Inline message identifier
// @param userID User identifier
func (client *Client) GetInlineGameHighScores(inlineMessageID string, userID int32) (*tdlib.GameHighScores, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":             "getInlineGameHighScores",
		"inline_message_id": inlineMessageID,
		"user_id":           userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores tdlib.GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}
