# Introduction
The go-accessoryhid package is an implementation of the [AOA 2.0](https://source.android.com/devices/accessories/aoa2) protocol for Android HID devices.

Dependence on [gousb](https://github.com/google/gousb) and [libusb-1.0](https://github.com/libusb/libusb/wiki).

# Documentation
TODO.

# Installation
- [gousb](https://github.com/google/gousb#dependencies)

```bash
go get -u github.com/Tryanks/go-accessoryhid
```

# Usage

Example:
```go
package main

import accessory "github.com/Tryanks/go-accessoryhid"

func main() {
	devices, err := accessory.GetDevices(2)
	if err != nil {
		panic(err)
	}
	phone := devices[0]
	defer phone.Close()
	keyboard, err := phone.Register(accessory.KeyboardReportDesc) // Register keyboard report descriptor
	err = keyboard.SendEvent([]byte{
		0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
	}) // Hid event (pressing 'a')
	if err != nil {
		panic(err)
	}
	err = keyboard.SendEvent([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}) // Hid event (releasing 'a')
	if err != nil {
		panic(err)
	}
}
```