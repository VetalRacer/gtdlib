// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// GetPassportAuthorizationFormAvailableElements Returns already available Telegram Passport elements suitable for completing a Telegram Passport authorization form. Result can be received only once for each authorization form
// @param autorizationFormID Authorization form identifier
// @param password Password of the current user
func (client *Client) GetPassportAuthorizationFormAvailableElements(autorizationFormID int32, password string) (*tdlib.PassportElementsWithErrors, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                "getPassportAuthorizationFormAvailableElements",
		"autorization_form_id": autorizationFormID,
		"password":             password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElementsWithErrors tdlib.PassportElementsWithErrors
	err = json.Unmarshal(result.Raw, &passportElementsWithErrors)
	return &passportElementsWithErrors, err

}
