// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
// @param chatID Chat identifier of the Invoice message
// @param messageID Message identifier
// @param orderInfoID Identifier returned by ValidateOrderInfo, or an empty string
// @param shippingOptionID Identifier of a chosen shipping option, if applicable
// @param credentials The credentials chosen by user for payment
func (client *Client) SendPaymentForm(chatID int64, messageID int64, orderInfoID string, shippingOptionID string, credentials tdlib.InputCredentials) (*tdlib.PaymentResult, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":              "sendPaymentForm",
		"chat_id":            chatID,
		"message_id":         messageID,
		"order_info_id":      orderInfoID,
		"shipping_option_id": shippingOptionID,
		"credentials":        credentials,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentResult tdlib.PaymentResult
	err = json.Unmarshal(result.Raw, &paymentResult)
	return &paymentResult, err

}
