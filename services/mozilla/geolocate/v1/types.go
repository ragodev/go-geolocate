package geolocate

type GeoRequest struct {
	// Internal stuff not sent to mozilla location services
	apiKey string `json:"-"`

	// The mobile country code (MCC) for the device's home network.
	HomeMobileCountryCode int `json:"homeMobileCountryCode,omitempty"`

	// The mobile network code (MNC) for the device's home network.
	HomeMobileNetworkCode int `json:"homeMobileNetworkCode,omitempty"`

	// The mobile radio type.
	RadioType string `json:"radioType,omitempty"`

	// The carrier name.
	Carrier string `json:"carrier,omitempty"`

	// Array of cell towers
	CellTowers []CellTower `json:"cellTowers,omitempty"`

	// Array of wifi access points
	WifiAccessPoints []WifiAccessPoint `json:"wifiAccessPoints,omitempty"`
}

type CellTower struct {
	// Unique identifier of the cell.
	CellId int `json:"cellId"`

	// The Location Area Code (LAC) for GSM and WCDMAnetworks. The Network ID (NID) for CDMA networks.
	LocationAreaCode int `json:"locationAreaCode"`

	// The cell tower's Mobile Country Code (MCC).
	MobileCountryCode int `json:"mobileCountryCode"`

	// The cell tower's Mobile Network Code.
	MobileNetworkCode int `json:"mobileNetworkCode"`

	// The number of milliseconds since this cell was primary.
	Age int `json:"age,omitempty"`

	// Radio signal strength measured in dBm.
	SignalStrength int `json:"signalStrength,omitempty"`

	// The timing advance value.
	TimingAdvance int `json:"timingAdvance,omitempty"`
}

type WifiAccessPoint struct {
	// The MAC address of the WiFi node.
	MacAddress string `json:"macAddress"`

	// The current signal strength measured in dBm.
	SignalStrength int `json:"signalStrength,omitempty"`

	// The number of milliseconds since this access point was detected.
	Age int `json:"age,omitempty"`

	// The channel over which the client is communicating with the access point.
	Channel int `json:"channel,omitempty"`

	// The current signal to noise ratio measured in dB.
	SignalToNoiseRatio int `json:"signalToNoiseRatio,omitempty"`
}

type GeoResponse struct {
	// The user’s estimated latitude and longitude, in degrees.
	Location Location `json:"location"`

	// The accuracy of the estimated location, in meters.
	Accuracy float32 `json:"accuracy"`
}

type Location struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type GeoResponseError struct {
	Error struct {
		// A list of errors which occurred. Each error contains an identifier for the type of error (the reason) and a short description (the message).
		Errors []struct {
			Domain       string `json:"domain"`
			Reason       string `json:"reason"`
			Message      string `json:"message"`
			ExtendedHelp string `json:"extendedHelp"`
		} `json:"errors"`

		// This is the same as the HTTP status of the response.
		Code int `json:"code"`

		// A short description of the error.
		Message string `json:"message"`
	} `json:"error"`
}
