// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogVerbosityLevel() (*tdlib.LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getLogVerbosityLevel",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel tdlib.LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) GetLogTagVerbosityLevel(tag string) (*tdlib.LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getLogTagVerbosityLevel",
		"tag":   tag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel tdlib.LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}
