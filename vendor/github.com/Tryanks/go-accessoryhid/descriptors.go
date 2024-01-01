package accessory

// MouseReportDesc is the HID report descriptor for a mouse.
var MouseReportDesc = []byte{
	0x05, 0x01, // Usage Page (Generic Desktop)
	0x09, 0x02, // Usage (Mouse)
	0xA1, 0x01, // Collection (Application)
	0x09, 0x01, //   Usage (Pointer)
	0xA1, 0x00, //   Collection (Physical)
	0x05, 0x09, //     Usage Page (Buttons)
	0x19, 0x01, //     Usage Minimum (1)
	0x29, 0x05, //     Usage Maximum (5)
	0x15, 0x00, //     Logical Minimum (0)
	0x25, 0x01, //     Logical Maximum (1)
	0x95, 0x05, //     Report Count (5)
	0x75, 0x01, //     Report Size (1)
	0x81, 0x02, //     Input (Data, Variable, Absolute): 5 buttons bits
	0x95, 0x01, //     Report Count (1)
	0x75, 0x03, //     Report Size (3)
	0x81, 0x01, //     Input (Constant): 3 bits padding
	0x05, 0x01, //     Usage Page (Generic Desktop)
	0x09, 0x30, //     Usage (X)
	0x09, 0x31, //     Usage (Y)
	0x09, 0x38, //     Usage (Wheel)
	0x15, 0x81, //     Local Minimum (-127)
	0x25, 0x7F, //     Local Maximum (127)
	0x75, 0x08, //     Report Size (8)
	0x95, 0x03, //     Report Count (3)
	0x81, 0x06, //     Input (Data, Variable, Relative): 3 position bytes (X, Y, Wheel)
	0xC0, //         End Collection
	0xC0, //       End Collection
}

const scHidKeyboardKeys = 102
const reportKeyboardMaxKeys = 6

// KeyboardReportDesc is the HID report descriptor for a keyboard.
var KeyboardReportDesc = []byte{
	0x05, 0x01, // Usage Page (Generic Desktop)
	0x09, 0x06, // Usage (Keyboard)
	0xA1, 0x01, // Collection (Application)
	0x05, 0x07, //   Usage Page (Key Codes)
	0x19, 0xE0, //   Usage Minimum (224)
	0x29, 0xE7, //   Usage Maximum (231)
	0x15, 0x00, //   Logical Minimum (0)
	0x25, 0x01, //   Logical Maximum (1)
	0x75, 0x01, //   Report Size (1)
	0x95, 0x08, //   Report Count (8)
	0x81, 0x02, //   Input (Data, Variable, Absolute): Modifier byte
	0x75, 0x08, //   Report Size (8)
	0x95, 0x01, //   Report Count (1)
	0x81, 0x01, //   Input (Constant): Reserved byte
	0x05, 0x08, //   Usage Page (LEDs)
	0x19, 0x01, //   Usage Minimum (1)
	0x29, 0x05, //   Usage Maximum (5)
	0x75, 0x01, //   Report Size (1)
	0x95, 0x05, //   Report Count (5)
	0x91, 0x02, //   Output (Data, Variable, Absolute): LED report
	0x75, 0x03, //   Report Size (3)
	0x95, 0x01, //   Report Count (1)
	0x91, 0x01, //   Output (Constant): LED report padding
	0x05, 0x07, //   Usage Page (Key Codes)
	0x19, 0x00, //   Usage Minimum (0)
	0x29, scHidKeyboardKeys - 1, //   Usage Maximum (101)
	0x15, 0x00, //   Logical Minimum (0)
	0x25, scHidKeyboardKeys - 1, //   Logical Maximum(101)
	0x75, 0x08, //   Report Size (8)
	0x95, reportKeyboardMaxKeys, //   Report Count (6)
	0x81, 0x00, //   Input (Data, Array): Keys
	0xC0, //       End Collection
}

// TouchscreenReportDesc is the HID report descriptor for a touchscreen.
var TouchscreenReportDesc = []byte{
	0x05, 0x0d, // USAGE_PAGE (Digitizers)
	0x09, 0x04, // USAGE (Touch Screen)
	0xa1, 0x01, // COLLECTION (Application)
	0x09, 0x22, //   USAGE (Finger)
	0xa1, 0x00, //   COLLECTION (Physical)
	0x09, 0x42, //     USAGE (Tip Switch)
	0x09, 0x51, //     USAGE (Contact Identifier) *added line
	0x15, 0x00, //     LOGICAL_MINIMUM (0)
	0x25, 0x01, //     LOGICAL_MAXIMUM (1)
	0x75, 0x01, //     REPORT_SIZE (1)
	0x95, 0x02, //     REPORT_COUNT (2)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0x95, 0x0e, //     REPORT_COUNT (14)
	0x81, 0x03, //     INPUT (Cnst,Var,Abs)
	0x05, 0x01, //     USAGE_PAGE (Generic Desktop)
	0x75, 0x10, //     REPORT_SIZE (16)
	0x95, 0x01, //     REPORT_COUNT (1)
	0x55, 0x0d, //     UNIT_EXPONENT (-3)
	0x65, 0x33, //     UNIT (Inch,EngLinear)
	0x15, 0x00, //     LOGICAL_MINIMUM (0)
	0x26, 0xff, 0x7f, //     LOGICAL_MAXIMUM (32767)
	0x09, 0x30, //     USAGE (X)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0x09, 0x31, //     USAGE (Y)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0xc0, //   END_COLLECTION
	0xc0, // END_COLLECTION
}

// ConsumerReportDesc is the HID report descriptor for a consumer control.
var ConsumerReportDesc = []byte{
	0x05, 0x0c, //        USAGE_PAGE (Consumer Devices)
	0x09, 0x01, //        USAGE (Consumer Control)
	0xa1, 0x01, //        COLLECTION (Application)
	0x15, 0x00, //          LOGICAL_MINIMUM (0)
	0x25, 0x01, //          LOGICAL_MAXIMUM (1)
	0x75, 0x01, //          REPORT_SIZE (1)
	0x95, 0x04, //          REPORT_COUNT (4)
	0x0a, 0x23, 0x02, //    USAGE (AC Home)
	0x0a, 0x24, 0x02, //    USAGE (AC Back)
	0x09, 0x40, //          USAGE (Menu)
	0x09, 0x34, //          USAGE (Power)
	0x81, 0x06, //          INPUT (Data,Var,Rel)
	0x75, 0x04, //          REPORT_SIZE (4)
	0x95, 0x01, //          REPORT_COUNT (1)
	0x81, 0x03, //          INPUT (Cnst,Var,Abs)
	0xc0, //              END_COLLECTION
}
