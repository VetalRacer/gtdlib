// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// SendCustomRequest Sends a custom request; for bots only
// @param method The method name
// @param parameters JSON-serialized method parameters
func (client *Client) SendCustomRequest(method string, parameters string) (*tdlib.CustomRequestResult, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "sendCustomRequest",
		"method":     method,
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var customRequestResult tdlib.CustomRequestResult
	err = json.Unmarshal(result.Raw, &customRequestResult)
	return &customRequestResult, err

}
