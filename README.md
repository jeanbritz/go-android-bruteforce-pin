# Android PIN Bruteforcer

Purpose of this project is to demonstrate the use of the [Android Open Accessory Protocol (AOA)](https://source.android.com/docs/core/interaction/accessories/protocol)
to mimic key presses on an Android device. The script goes and try to type each PIN in the Android device
in the list of pins located [here](/pins) 

The only requirements to use this project, are the following:
* Android device with the accessory mode capability
* USB Cable
* Linux OS with Go 1.22.0 installed

# Prerequisites
If you are running this project for the first time on Linux, then you should install `libusb` as well,
if you have not already. You can install it as follows:
```shell
$ make install-libusb
```
or, alternatively:
```shell
# apt-get install libusb-1.0-0-dev
```

# Usage
To run the script, first build the project if you haven't yet:
```shell
$ make build
```

Then run it with sudo privileges:
```shell
$ make run
```

# Details
I have used a [Sony Xperia Z1 Compact](https://www.gsmarena.com/sony_xperia_z1_compact-5753.php) as the Android phone, which I am still trying to bruteforce the PIN.

This script registers itself as an accessory with the Android phone, which allows me to interact with HID Events over USB.
The Android OAO Protocol only sets up the initial part of registering the script as an accessory.
This setup includes a [HID report descriptor](pkg/hid/descriptor.go), which is sent to the Android phone to configure it.
What this report descriptor does is it configures a mouse pointer and a stylus on the Android phone.
For my case it made it easier for me to have both of them as the mouse pointer allowed me to use absolute values to position the pointer
and then convert it to the stylus to emulate a touch to press any of the 10 digits (0-9)

For each HID event sent to the Android phone you need to send 5 bytes (depends on the report descriptor)
Format for each action performed on the phone is described below:

| Action                                            | Byte 0 | Byte 1      | Byte 2      | Byte 3      | Byte 4      |
|---------------------------------------------------|--------|-------------|-------------|-------------|-------------|
| Press                                             | 0x00   | 0x00        | 0x00        | 0x00        | 0x00        |
| Convert Pointer <br/>to Stylus accessory          | 0x01   | 0x00        | 0x00        | 0x00        | 0x00        | 
| Convert to Pointer and <br/> Set Pointer Position | 0x02   | X Pos (LSB) | X Pos (MSB) | Y Pos (LSB) | Y Pos (MSB) |

# Problems
There are a few problems with this script. Not due to bugs, I think, but due to the protocol being asynchronous.
There is no feedback or error handling that can be implemented when send a HID event.

Through testing the script I have found that when sending the HID event, it does not always mean the action 
(described above) is executed 100% reliably everytime.

This means that you will need to 'babysit' the phone to see if every PIN combination is entered correctly.

# Performance
It will vary between devices due to the backoff period between many PIN entries.
On my Sony you have to wait for 30 seconds after you have entered 5 PINs.

For the Sony it takes on average about 8 seconds to enter 5 pins. Then 30 second of backoff time to wait before you
can enter another 5 pins.

# Demo
![Demo](media/android-bruteforce-demo.webp)

# Credits
* Thanks to [Tryanks](https://github.com/Tryanks) for posting [this repository](https://github.com/Tryanks/go-accessoryhid)
* Thanks to [urbanadventurer](https://github.com/urbanadventurer) for posting a list of [optimized PIN combinations](https://github.com/urbanadventurer/Android-PIN-Bruteforce)

# Useful tools
* [USB Descriptor and Request Parser](https://eleccelerator.com/usbdescreqparser/)

# Related links
* https://www.beyondlogic.org/usbnutshell/usb6.shtml