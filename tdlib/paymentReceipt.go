// AUTOGENERATED - DO NOT EDIT

package tdlib

// PaymentReceipt Contains information about a successful payment
type PaymentReceipt struct {
	tdCommon
	Date                   int32           `json:"date"`                      // Point in time (Unix timestamp) when the payment was made
	PaymentsProviderUserID int32           `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	Invoice                *Invoice        `json:"invoice"`                   // Contains information about the invoice
	OrderInfo              *OrderInfo      `json:"order_info"`                // Contains order information; may be null
	ShippingOption         *ShippingOption `json:"shipping_option"`           // Chosen shipping option; may be null
	CredentialsTitle       string          `json:"credentials_title"`         // Title of the saved credentials
}

// MessageType return the string telegram-type of PaymentReceipt
func (paymentReceipt *PaymentReceipt) MessageType() string {
	return "paymentReceipt"
}

// NewPaymentReceipt creates a new PaymentReceipt
//
// @param date Point in time (Unix timestamp) when the payment was made
// @param paymentsProviderUserID User identifier of the payment provider bot
// @param invoice Contains information about the invoice
// @param orderInfo Contains order information; may be null
// @param shippingOption Chosen shipping option; may be null
// @param credentialsTitle Title of the saved credentials
func NewPaymentReceipt(date int32, paymentsProviderUserID int32, invoice *Invoice, orderInfo *OrderInfo, shippingOption *ShippingOption, credentialsTitle string) *PaymentReceipt {
	paymentReceiptTemp := PaymentReceipt{
		tdCommon:               tdCommon{Type: "paymentReceipt"},
		Date:                   date,
		PaymentsProviderUserID: paymentsProviderUserID,
		Invoice:                invoice,
		OrderInfo:              orderInfo,
		ShippingOption:         shippingOption,
		CredentialsTitle:       credentialsTitle,
	}

	return &paymentReceiptTemp
}
