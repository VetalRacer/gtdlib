// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline request if only_local is true. Can be called before authorization
// @param onlyLocal If true, returns only locally available information without sending network requests
func (client *Client) GetLocalizationTargetInfo(onlyLocal bool) (*tdlib.LocalizationTargetInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getLocalizationTargetInfo",
		"only_local": onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var localizationTargetInfo tdlib.LocalizationTargetInfo
	err = json.Unmarshal(result.Raw, &localizationTargetInfo)
	return &localizationTargetInfo, err

}
