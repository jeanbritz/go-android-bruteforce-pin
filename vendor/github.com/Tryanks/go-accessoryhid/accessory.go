package accessory

type Accessory struct {
	hidID        uint16
	ParentDevice *AccessoryDevice
}

func newAccessory(hidID uint16, device *AccessoryDevice) *Accessory {
	return &Accessory{
		hidID:        hidID,
		ParentDevice: device,
	}
}

func (a *Accessory) SendEvent(event []byte) error {
	return a.ParentDevice.SendHidEvent(a.hidID, event)
}

func (a *Accessory) Unregister() error {
	return a.ParentDevice.Unregister(a.hidID)
}
