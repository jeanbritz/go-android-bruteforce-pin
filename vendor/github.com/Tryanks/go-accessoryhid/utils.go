package accessory

import (
	"encoding/binary"
	"fmt"
	"github.com/google/gousb"
	"math/rand"
	"time"
)

// SkipList is a list of vendor IDs that are known to not support the accessory protocol
// and should be skipped
// TODO: Add more vendor IDs to the list
var SkipList = []uint16{
	0x8087, // Intel Corp.
	0x1d6b, // Linux Foundation
	0x2109, // VIA Labs, Inc.
}

// GetDevices return a list of devices that support the specified protocol version
func GetDevices(protocolVersion uint16) (accessoryList []*AccessoryDevice, err error) {
	accessoryList = make([]*AccessoryDevice, 0)
	devices, err := gousb.NewContext().OpenDevices(func(desc *gousb.DeviceDesc) bool {
		for _, id := range SkipList {
			if desc.Vendor == gousb.ID(id) {
				return false
			}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	for _, d := range devices {
		p, err := getProtocol(d)
		if err != nil || p < protocolVersion {
			d.Close()
			fmt.Println("No Protocol!")
			continue
		}
		manu, err := d.Manufacturer()
		if err != nil {
			d.Close()
			fmt.Println("No ManuFacturer!")
			continue
		}
		accessoryList = append(accessoryList, NewAccessoryDevice(d, p, manu))
	}
	return accessoryList, nil
}

// getProtocol return the protocol version of the device
func getProtocol(dev *gousb.Device) (protocol uint16, err error) {
	dev.ControlTimeout = 500 * time.Millisecond
	if dev == nil {
		return 0, ErrorNoAccessoryDevice
	}
	data := make([]byte, 2)
	n, err := dev.Control(RTypeIn, AccessoryGetProtocol, 0, 0, data)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, ErrorFailedToGetProtocol
	}
	protocol = binary.LittleEndian.Uint16(data)
	return
}

// uint16InList return the index of the value in the list
func uint16InList(list []uint16, value uint16) (int, bool) {
	for i, v := range list {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

// uint16GetUniqueRandom return a random uint16 that is not in the list
func uint16GetUniqueRandom(list []uint16) uint16 {
	for {
		r := uint16GetRandom()
		if r <= 100 {
			continue
		}
		if _, ok := uint16InList(list, r); !ok {
			return r
		}
	}
}

// uint16GetRandom return a random uint16
func uint16GetRandom() uint16 {
	return uint16(rand.Intn(0xffff))
}
