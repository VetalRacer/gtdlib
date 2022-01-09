// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization
// @param onlyCurrent If true, returns only data for the current library launch
func (client *Client) GetNetworkStatistics(onlyCurrent bool) (*tdlib.NetworkStatistics, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "getNetworkStatistics",
		"only_current": onlyCurrent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var networkStatistics tdlib.NetworkStatistics
	err = json.Unmarshal(result.Raw, &networkStatistics)
	return &networkStatistics, err

}
