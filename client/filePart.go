// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
// @param fileID Identifier of the file. The file must be located in the TDLib file cache
// @param offset The offset from which to read the file
// @param count Number of bytes to read. An error will be returned if there are not enough bytes available in the file from the specified position. Pass 0 to read all available data from the specified position
func (client *Client) ReadFilePart(fileID int32, offset int32, count int32) (*tdlib.FilePart, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "readFilePart",
		"file_id": fileID,
		"offset":  offset,
		"count":   count,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart tdlib.FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err

}
