// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/vetalracer/gtdlib/tdlib"
)

// GetLogTags Returns list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (client *Client) GetLogTags() (*tdlib.LogTags, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getLogTags",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logTags tdlib.LogTags
	err = json.Unmarshal(result.Raw, &logTags)
	return &logTags, err

}
