package geolocate

import "testing"

const TestApiKey = "test"

func TestGetCurrentLocation(t *testing.T) {

	greq := NewGeoRequest(TestApiKey)
	greq.HomeMobileCountryCode = 310
	greq.HomeMobileNetworkCode = 410
	greq.RadioType = "gsm"
	greq.Carrier = "Vodafone"
	greq.CellTowers = []CellTower{
		// GSM
		CellTower{

			CellId:            42,
			LocationAreaCode:  415,
			MobileCountryCode: 310,
			MobileNetworkCode: 410,
			Age:               0,
			SignalStrength:    -60,
			TimingAdvance:     15,
		},
		// WCDMA
		CellTower{

			CellId:            21532831,
			LocationAreaCode:  2862,
			MobileCountryCode: 214,
			MobileNetworkCode: 7,
		},
	}
	greq.WifiAccessPoints = []WifiAccessPoint{
		WifiAccessPoint{

			MacAddress:         "01:23:45:67:89:AB",
			SignalStrength:     -65,
			Age:                0,
			Channel:            11,
			SignalToNoiseRatio: 40,
		},
		WifiAccessPoint{

			MacAddress:     "01:23:45:67:89:AC",
			SignalStrength: 4,
			Age:            0,
		},
	}

	gresp, err := greq.GetCurrentLocation()
	if err != nil {

		t.Fatal(err)
	}

	t.Logf("Location Lat: %v\n", gresp.Location.Lat)
	t.Logf("Location Lng: %v\n", gresp.Location.Lng)
	t.Logf("Location Accuracy: %v\n", gresp.Accuracy)
}
