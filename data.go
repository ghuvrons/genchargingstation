package genchargingstation

import "fmt"

type ChargingData struct {
	StationID string
	BatteryID string
	Volt      float32
	Curr      float32
	Temp      float32
	Capc      float32
	SOC       float32
	SOH       float32
	Cycle     int
	Error     uint8
}

func (data *ChargingData) ToString() string {
	var g string = fmt.Sprintf(`
==== Charging Data ====
StationID : %s
BatteryID : %s
Volt      : %f
Curr      : %f
Temp      : %f
Capc      : %f
SOC       : %f
SOH       : %f
Cycle     : %d
Error     : %d
	`, data.StationID, data.BatteryID, data.Volt, data.Curr, data.Temp, data.Capc, data.SOC, data.SOH,
		data.Cycle, data.Error)
	return g
}
