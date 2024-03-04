package hid

import (
	"fmt"
	accessory "github.com/Tryanks/go-accessoryhid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/utils"
)

type Touchscreen struct {
	Accessory *accessory.Accessory
}

type Actions interface {
	SetPosition(x int16, y int16)
	Press()
}

func (t *Touchscreen) SetPosition(x int16, y int16) {
	xMsb := utils.GetMSB(x)
	xLsb := utils.GetLSB(x)
	yMsb := utils.GetMSB(y)
	yLsb := utils.GetLSB(y)

	// Use Pointer to set absolute coordinates
	err := t.Accessory.SendEvent([]byte{
		0x02, byte(xLsb), byte(xMsb), byte(yLsb), byte(yMsb),
	})

	if err != nil {
		err = fmt.Errorf("error occurred while setting pointer position %w", err)
		fmt.Println(err)
	}
}

func (t *Touchscreen) Press() {
	// Convert Pointer to Touch Accessory
	err := t.Accessory.SendEvent([]byte{
		0x01, 0, 0, 0, 0,
	})

	if err != nil {
		err = fmt.Errorf("error occurred while converting %w", err)
		fmt.Println(err)
	}
	// Press
	err = t.Accessory.SendEvent([]byte{
		0x00, 0, 0, 0, 0,
	})

	if err != nil {
		err = fmt.Errorf("error occurred while trying to press %w", err)
		fmt.Println(err)
	}
}
