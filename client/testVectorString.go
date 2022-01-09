// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of strings to return
func (client *Client) TestCallVectorString(x []string) (*tdlib.TestVectorString, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testCallVectorString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorString tdlib.TestVectorString
	err = json.Unmarshal(result.Raw, &testVectorString)
	return &testVectorString, err

}
