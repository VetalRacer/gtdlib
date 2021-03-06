// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked; should begin with "https://t.me/joinchat/", "https://telegram.me/joinchat/", or "https://telegram.dog/joinchat/"
func (client *Client) CheckChatInviteLink(inviteLink string) (*tdlib.ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkInfo tdlib.ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err

}
