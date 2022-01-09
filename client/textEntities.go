// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) contained in the text. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*tdlib.TextEntities, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var textEntities tdlib.TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err

}
