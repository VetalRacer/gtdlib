// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// CreateTemporaryPassword Creates a new temporary password for processing payments
// @param password Persistent user password
// @param validFor Time during which the temporary password will be valid, in seconds; should be between 60 and 86400
func (client *Client) CreateTemporaryPassword(password string, validFor int32) (*tdlib.TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "createTemporaryPassword",
		"password":  password,
		"valid_for": validFor,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState tdlib.TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}

// GetTemporaryPasswordState Returns information about the current temporary password
func (client *Client) GetTemporaryPasswordState() (*tdlib.TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getTemporaryPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState tdlib.TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}