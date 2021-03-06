// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetRecentlyVisitedTMeURLs Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeURLs(referrer string) (*tdlib.TMeURLs, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tMeURLs tdlib.TMeURLs
	err = json.Unmarshal(result.Raw, &tMeURLs)
	return &tMeURLs, err

}
