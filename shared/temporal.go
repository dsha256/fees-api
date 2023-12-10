package shared

var SignalChannels = struct {
	OpenBillChannel      string
	GetBillByUUIDChannel string
	AddLineItem          string
}{
	OpenBillChannel:      "OPEN_BILL_CHANNEL",
	GetBillByUUIDChannel: "GET_BILL_BY_UUID_CHANNEL",
	AddLineItem:          "ADD_LINE_ITEM_CHANNEL",
}

type (
	OpenBillSignal struct {
		UUID string
	}
)
