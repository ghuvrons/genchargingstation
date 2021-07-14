package genchargingstation

import "fmt"

type ChargingData struct {
	StationID string
	Hole      uint8
	BatteryID uint32
	Volt      float32
	Curr      float32
	Temp      float32
	Capc      float32
	SOC       int
	SOH       int
	Cycle     int
	Error     uint8
	LocLong   float64
	LocLat    float64
}

func (data *ChargingData) ToString() string {
	var g string = fmt.Sprintf(`
==== Charging Data ====
StationID : %s
BatteryID : %d
Volt      : %f
Curr      : %f
Temp      : %f
Capc      : %f
SOC       : %d
SOH       : %d
Cycle     : %d
Error     : %d
	`, data.StationID, data.BatteryID, data.Volt, data.Curr, data.Temp, data.Capc, data.SOC, data.SOH,
		data.Cycle, data.Error)
	return g
}
