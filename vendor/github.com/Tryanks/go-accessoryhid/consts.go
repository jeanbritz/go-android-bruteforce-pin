package accessory

import (
	"github.com/google/gousb"
)

// AOAv2 Basic Protocol
const (
	AccessoryRegisterHid      uint8 = 54
	AccessoryUnregisterHid          = 55
	AccessorySetHidReportDesc       = 56
	AccessorySendHidEvent           = 57

	AccessoryGetProtocol = 51
)

// USB Control transfer type
const (
	RTypeOut uint8 = gousb.ControlOut | gousb.ControlVendor
	RTypeIn        = gousb.ControlIn | gousb.ControlVendor
)
