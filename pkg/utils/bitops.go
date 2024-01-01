package utils

func GetLSB(value int16) int {
	return int(value & 0xff)
}

func GetMSB(value int16) int {
	return int((value >> 8) & 0xff)
}
