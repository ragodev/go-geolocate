// Package geolocate provides an implementation of the Mozilla Geo Location Service API.
package geolocate

import (
	"fmt"
	"net"

	"github.com/jmcvetta/napping"
)

// The URL of the Mozilla Geo Location API.
const BaseURL = "https://location.services.mozilla.com/v1/geolocate"

// NewGeoRequest creates a new GeoRequest object and takes a
// API key.
func NewGeoRequest(key string) (greq *GeoRequest) {
	return &GeoRequest{apiKey: key}
}

// AddGSMTower adds a new GSM tower to the GeoRequest object and performs some sanity checks.
func (greq *GeoRequest) AddGSMTower(cellId,
	locationAreaCode,
	mobileCountryCode,
	mobileNetworkCode,
	age,
	signalStrength,
	timingAdvance int) (err error) {
	greq.CellTowers = append(greq.CellTowers, CellTower{
		CellId:            cellId,
		LocationAreaCode:  locationAreaCode,
		MobileCountryCode: mobileCountryCode,
		MobileNetworkCode: mobileNetworkCode,
		Age:               age,
		SignalStrength:    signalStrength,
		TimingAdvance:     timingAdvance,
	})
	return
}

// AddWCDMATower adds a new WCDMA tower to the GeoRequest object and performs some sanity checks.
func (greq *GeoRequest) AddWCDMATower(cellId,
	locationAreaCode,
	mobileCountryCode,
	mobileNetworkCode int) (err error) {
	greq.CellTowers = append(greq.CellTowers, CellTower{
		CellId:            cellId,
		LocationAreaCode:  locationAreaCode,
		MobileCountryCode: mobileCountryCode,
		MobileNetworkCode: mobileNetworkCode,
	})
	return
}

// AddWifiAccessPoint adds a new Wifi access point to the GeoRequest object and performs
// some sanity checks.
func (greq *GeoRequest) AddWifiAccessPoint(macAddress string,
	signalStrength,
	age,
	channel,
	signalToNoiseRatio int) (err error) {
	_, err = net.ParseMAC(macAddress)
	if err != nil {
		return
	}

	greq.WifiAccessPoints = append(greq.WifiAccessPoints, WifiAccessPoint{
		MacAddress:         macAddress,
		SignalStrength:     signalStrength,
		Age:                age,
		Channel:            channel,
		SignalToNoiseRatio: signalToNoiseRatio,
	})
	return
}

// GetCurrentLocation sends everything in the GeoRequest object to the Mozilla Geo Location service
// and returns the parsed response.
func (greq *GeoRequest) GetCurrentLocation() (gresp *GeoResponse, err error) {

	// Perform some validation checks
	if (len(greq.CellTowers) == 0) && (len(greq.WifiAccessPoints) == 0) {
		return nil, fmt.Errorf("No cell towers or wifi access points were provided.")
	}

	gresp = new(GeoResponse)
	gerr := new(GeoResponseError)
	req := &napping.Request{
		Url:     BaseURL,
		Method:  "Post",
		Payload: greq,
		Result:  gresp,
		Error:   gerr,
	}
	if len(greq.apiKey) != 0 {
		req.Params = &napping.Params{"key": greq.apiKey}
	}

	resp, err := napping.Send(req)
	if err != nil {
		return
	}
	if resp.Status() >= 400 {
		return nil, fmt.Errorf("Bad response from Mozilla Geo location API. Message: %s", gerr.Error.Message)
	}
	return
}
