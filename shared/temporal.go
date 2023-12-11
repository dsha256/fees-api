package shared

import (
	"github.com/dsha256/feesapi/currency"
)

var SignalChannels = struct {
	OpenBillChannel      string
	CheckoutBillChannel  string
	GetBillByUUIDChannel string
	AddLineItem          string
}{
	OpenBillChannel:      "OPEN_BILL_CHANNEL",
	CheckoutBillChannel:  "CHECKOUT_CHANNEL",
	GetBillByUUIDChannel: "GET_BILL_BY_UUID_CHANNEL",
	AddLineItem:          "ADD_LINE_ITEM_CHANNEL",
}

type (
	OpenBillSignal struct {
		UUID string
	}

	CheckoutSignal struct {
		PayingCurrency currency.SupportedCurrency
	}
)
