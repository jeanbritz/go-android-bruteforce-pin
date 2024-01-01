package accessory

import "github.com/google/gousb"

// AccessoryDevice Connected Android Device
type AccessoryDevice struct {
	Device       *gousb.Device
	Protocol     uint16
	Manufacturer string
	hidList      []uint16
}

// NewAccessoryDevice Create a new AccessoryDevice
func NewAccessoryDevice(device *gousb.Device, protocol uint16, manu string) *AccessoryDevice {
	return &AccessoryDevice{
		Device:       device,
		Protocol:     protocol,
		Manufacturer: manu,
		hidList:      make([]uint16, 0),
	}
}

func (a *AccessoryDevice) SendHidEvent(hidID uint16, event []byte) error {
	_, err := a.Device.Control(RTypeOut, AccessorySendHidEvent, hidID, 0, event)
	return err
}

func (a *AccessoryDevice) Register(reportDesc []byte) (accessory *Accessory, err error) {
	hidID := uint16GetUniqueRandom(a.hidList)
	_, err = a.Device.Control(RTypeOut, AccessoryRegisterHid, hidID, uint16(len(reportDesc)), nil)
	if err != nil {
		return
	}
	_, err = a.Device.Control(RTypeOut, AccessorySetHidReportDesc, hidID, 0, reportDesc)
	if err != nil {
		return
	}
	a.hidList = append(a.hidList, hidID)
	accessory = newAccessory(hidID, a)
	return
}

func (a *AccessoryDevice) Unregister(hidID uint16) error {
	if i, ok := uint16InList(a.hidList, hidID); ok {
		a.hidList = append(a.hidList[:i], a.hidList[i+1:]...)
	}
	_, err := a.Device.Control(RTypeOut, AccessoryUnregisterHid, hidID, 0, nil)
	return err
}

func (a *AccessoryDevice) Close() error {
	for _, hidID := range a.hidList {
		_ = a.Unregister(hidID)
	}
	return a.Device.Close()
}
