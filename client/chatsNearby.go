// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// SearchChatsNearby Returns a list of users and location-based supergroups nearby. The list of users nearby will be updated for 60 seconds after the request by the updates updateUsersNearby. The request should be sent again every 25 seconds with adjusted location to not miss new chats
// @param location Current user location
func (client *Client) SearchChatsNearby(location *tdlib.Location) (*tdlib.ChatsNearby, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "searchChatsNearby",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatsNearby tdlib.ChatsNearby
	err = json.Unmarshal(result.Raw, &chatsNearby)
	return &chatsNearby, err

}
