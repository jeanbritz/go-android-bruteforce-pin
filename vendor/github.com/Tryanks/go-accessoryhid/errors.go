package accessory

import "errors"

// Error definitions
var (
	ErrorNoAccessoryDevice   = errors.New("ErrorNoAccessoryDevice")
	ErrorFailedToGetProtocol = errors.New("ErrorFailedToGetProtocol")
)
